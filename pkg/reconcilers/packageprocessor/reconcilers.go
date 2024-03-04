/*
Copyright 2024 Nokia.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package packageprocessor

import (
	"context"
	"fmt"
	"path"
	"strings"

	"github.com/henderiw/logger/log"
	"github.com/henderiw/resource"
	"github.com/henderiw/store"
	"github.com/henderiw/store/memory"
	"github.com/kform-dev/kform/pkg/kform/runner"
	"github.com/kform-dev/kform/pkg/pkgio"
	"github.com/kform-dev/pkg-server/apis/condition"
	pkgv1alpha1 "github.com/kform-dev/pkg-server/apis/pkg/v1alpha1"
	"github.com/kform-dev/pkg-server/apis/pkgid"
	"github.com/kform-dev/pkg-server/pkg/reconcilers"
	"github.com/kform-dev/pkg-server/pkg/reconcilers/ctrlconfig"
	"github.com/kform-dev/pkg-server/pkg/reconcilers/lease"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

func init() {
	reconcilers.Register("packageprocessor", &reconciler{})
}

const (
	controllerName = "PackageProcessorController"
	//controllerEventError = "PkgRevInstallerError"
	controllerEvent = "PackageProcessor"
	finalizer       = "packageprocessor.pkg.kform.dev/finalizer"
	// errors
	errGetCr            = "cannot get cr"
	errUpdateStatus     = "cannot update status"
	controllerCondition = condition.ReadinessGate_PkgProcess
)

//+kubebuilder:rbac:groups=pkg.kform.dev,resources=packagerevision,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=pkg.kform.dev,resources=packagerevision/status,verbs=get;update;patch

// SetupWithManager sets up the controller with the Manager.
func (r *reconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, c interface{}) (map[schema.GroupVersionKind]chan event.GenericEvent, error) {
	/*
		if err := configv1alpha1.AddToScheme(mgr.GetScheme()); err != nil {
			return nil, err
		}
	*/
	r.APIPatchingApplicator = resource.NewAPIPatchingApplicator(mgr.GetClient())
	r.finalizer = resource.NewAPIFinalizer(mgr.GetClient(), finalizer)
	r.recorder = mgr.GetEventRecorderFor(controllerName)

	return nil, ctrl.NewControllerManagedBy(mgr).
		Named(controllerName).
		For(&pkgv1alpha1.PackageRevision{}).
		Complete(r)
}

type reconciler struct {
	resource.APIPatchingApplicator
	finalizer *resource.APIFinalizer
	recorder  record.EventRecorder
}

func (r *reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	ctx = ctrlconfig.InitContext(ctx, controllerName, req.NamespacedName)
	log := log.FromContext(ctx)
	log.Info("reconcile")

	cr := &pkgv1alpha1.PackageRevision{}
	if err := r.Get(ctx, req.NamespacedName, cr); err != nil {
		// if the resource no longer exists the reconcile loop is done
		if resource.IgnoreNotFound(err) != nil {
			log.Error(errGetCr, "error", err)
			return ctrl.Result{}, errors.Wrap(resource.IgnoreNotFound(err), errGetCr)
		}
		return ctrl.Result{}, nil
	}
	key := types.NamespacedName{
		Namespace: cr.GetNamespace(),
		Name:      cr.GetName(),
	}
	cr = cr.DeepCopy()
	// if the pkgRev is a catalog packageRevision or
	// if the pkgrev does not have a condition to process the pkgRev
	// this event is not relevant
	if strings.HasPrefix(cr.GetName(), pkgid.PkgTarget_Catalog) ||
		!cr.HasReadinessGate(controllerCondition) {
		return ctrl.Result{}, nil
	}
	// the package revisioner first need to act e.g. to clone the repo
	if cr.GetCondition(condition.ConditionTypeReady).Status == metav1.ConditionFalse {
		return ctrl.Result{}, nil
	}
	// the previous condition is not finished, so we need to wait for acting
	if cr.GetPreviousCondition(controllerCondition).Status == metav1.ConditionFalse {
		return ctrl.Result{}, nil
	}

	l := lease.New(r.Client, key)
	if err := l.AcquireLease(ctx, controllerName); err != nil {
		log.Debug("cannot acquire lease", "key", key.String(), "error", err.Error())
		r.recorder.Eventf(cr, corev1.EventTypeWarning,
			"lease", "error %s", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: lease.RequeueInterval}, nil
	}
	r.recorder.Eventf(cr, corev1.EventTypeWarning,
		"lease", "acquired")

	if !cr.GetDeletionTimestamp().IsZero() {
		// TODO release resources that were allocated e.g. ipam?

		// remove the finalizer
		if err := r.finalizer.RemoveFinalizer(ctx, cr); err != nil {
			//log.Error("cannot remove finalizer", "error", err)
			r.recorder.Eventf(cr, corev1.EventTypeWarning,
				controllerEvent, "error %s", err.Error())
			return ctrl.Result{Requeue: true}, nil
		}
		return ctrl.Result{}, nil
	}

	if err := r.finalizer.AddFinalizer(ctx, cr); err != nil {
		log.Error("cannot add finalizer", "error", err)
		r.recorder.Eventf(cr, corev1.EventTypeWarning,
			controllerEvent, "error %s", err.Error())
		cr.SetConditions(condition.ConditionUpdate(controllerCondition, "cannot add finalizer", err.Error()))
		return ctrl.Result{Requeue: true}, errors.Wrap(r.Status().Update(ctx, cr), errUpdateStatus)
	}

	// process resource data
	// run kform
	// determine error
	// create output
	if cr.GetCondition(controllerCondition).Status == metav1.ConditionFalse {
		pkgrevResources := &pkgv1alpha1.PackageRevisionResources{}
		if err := r.Get(ctx, key, pkgrevResources); err != nil {
			log.Error("cannot get resources from pkgRevResources", "key", key, "error", err)
			r.recorder.Eventf(cr, corev1.EventTypeWarning,
				controllerEvent, "error %s", err.Error())
			cr.SetConditions(condition.ConditionUpdate(controllerCondition, "cannot get resources from pkgRevResources", err.Error()))
			return ctrl.Result{Requeue: true}, errors.Wrap(r.Status().Update(ctx, cr), errUpdateStatus)
		}
		resourceData := memory.NewStore[[]byte]()
		for k, v := range pkgrevResources.Spec.Resources {
			// Replace the ending \r\n (line ending used in windows) with \n and then split it into multiple YAML documents
			// if it contains document separators (---)
			values, err := pkgio.SplitDocuments(strings.ReplaceAll(v, "\r\n", "\n"))
			if err != nil {
				continue
			}
			for i := range values {
				// the Split used above will eat the tail '\n' from each resource. This may affect the
				// literal string value since '\n' is meaningful in it.
				if i != len(values)-1 {
					values[i] += "\n"
				}
				resourceData.Create(ctx, store.ToKey(fmt.Sprintf("%s.%d", k, i)), []byte(values[i]))
			}
			//resourceData.Create(ctx, store.ToKey(k), []byte(v))
		}
		// process input
		inputData := memory.NewStore[[]byte]()
		for i, input := range cr.Spec.Inputs {
			inputData.Create(ctx, store.ToKey(fmt.Sprintf("pkgctx-%d", i)), input.Raw)
		}

		outputData := memory.NewStore[[]byte]()
		kfrunner := runner.NewKformRunner(&runner.Config{
			PackageName:  cr.Spec.PackageID.Package,
			InputData:    inputData,
			ResourceData: resourceData,
			OutputData:   outputData,
		})

		if err := kfrunner.Run(ctx); err != nil {
			log.Error("kform run failed", "error", err)
			r.recorder.Eventf(cr, corev1.EventTypeWarning,
				controllerEvent, "error %s", err.Error())
			cr.SetConditions(condition.ConditionUpdate(controllerCondition, "kform run failed", err.Error()))
			return ctrl.Result{Requeue: true}, errors.Wrap(r.Status().Update(ctx, cr), errUpdateStatus)
		}

		resources := map[string]string{}
		outputData.List(ctx, func(ctx context.Context, key store.Key, data []byte) {
			fmt.Println("key", key.Name)
			//fmt.Println("data", string(data))
			resources[path.Join("out", key.Name)] = string(data)
		})
		newPkgrevResources := pkgv1alpha1.BuildPackageRevisionResources(
			*cr.ObjectMeta.DeepCopy(),
			pkgv1alpha1.PackageRevisionResourcesSpec{
				PackageID: *cr.Spec.PackageID.DeepCopy(),
				Resources: resources,
			},
			pkgv1alpha1.PackageRevisionResourcesStatus{},
		)
		//pkgrevResources.Spec.Resources = resources
		if err := r.Update(ctx, newPkgrevResources); err != nil {
			log.Error("cannot update pkgRevResources", "error", err)
			r.recorder.Eventf(cr, corev1.EventTypeWarning,
				controllerEvent, "error %s", err.Error())
			cr.SetConditions(condition.ConditionUpdate(controllerCondition, "cannot update pkgRevResources", err.Error()))
			return ctrl.Result{Requeue: true}, errors.Wrap(r.Status().Update(ctx, cr), errUpdateStatus)
		}

		log.Debug("process complete")
		r.recorder.Eventf(cr, corev1.EventTypeNormal,
			controllerEvent, "ready")
		cr.SetConditions(condition.ConditionReady(controllerCondition))
		cType := cr.NextReadinessGate(controllerCondition)
		if cType != condition.ConditionTypeEnd {
			if !cr.HasCondition(cType) {
				cr.SetConditions(condition.ConditionUpdate(cType, "init", ""))
			}
		}
		return ctrl.Result{}, errors.Wrap(r.Status().Update(ctx, cr), errUpdateStatus)
	}
	return ctrl.Result{}, nil
}