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

package packagerevisionresource

import (
	"context"
	"fmt"
	"reflect"
	"sort"

	"github.com/henderiw/logger/log"
	"github.com/kform-dev/pkg-server/apis/condition"
	pkgv1alpha1 "github.com/kform-dev/pkg-server/apis/pkg/v1alpha1"
	"github.com/kform-dev/pkg-server/apis/pkgid"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *strategy) List(ctx context.Context, options *metainternalversion.ListOptions) (runtime.Object, error) {
	log := log.FromContext(ctx)
	log.Info("list")
	filter, err := parsePackageFieldSelector(ctx, options.FieldSelector)
	if err != nil {
		return nil, err
	}

	opts := []client.ListOption{}
	if options.LabelSelector != nil {
		opts = append(opts, client.MatchingLabelsSelector{Selector: options.LabelSelector})
	}

	//	if options.FieldSelector != nil {
	//		requirements := options.FieldSelector.Requirements()
	//		for _, req := range requirements {
	//			log.Info("list field selector requirements", "operator", req.Operator, "field", req.Field, "value", req.Value)
	//			opts = append(opts, client.MatchingFields{req.Field: req.Value})
	//		}
	//	}

	pkgRevs := pkgv1alpha1.PackageRevisionList{}
	if err := r.client.List(ctx, &pkgRevs, opts...); err != nil {
		return nil, apierrors.NewInternalError(err)
	}

	newListObj := &pkgv1alpha1.PackageRevisionResourcesList{}
	v, err := getListPrt(newListObj)
	if err != nil {
		return nil, err
	}

	for _, pkgRev := range pkgRevs.Items {
		if pkgRev.GetCondition(condition.ConditionTypeReady).Status == metav1.ConditionFalse {
			// the package might not be cloned yet
			continue
		}
		if filter != nil {
			if filter.Name != "" && filter.Name != pkgRev.Name {
				continue
			}
			if filter.Namespace != "" && filter.Namespace != pkgRev.Namespace {
				continue
			}
			if filter.Target != "" && filter.Target != pkgRev.Spec.PackageID.Target {
				continue
			}
			if filter.Repository != "" && filter.Repository != pkgRev.Spec.PackageID.Repository {
				continue
			}
			if filter.Realm != "" && filter.Realm != pkgRev.Spec.PackageID.Realm {
				continue
			}
			if filter.Package != "" && filter.Package != pkgRev.Spec.PackageID.Package {
				continue
			}
			if filter.Revision != "" && filter.Revision != pkgRev.Spec.PackageID.Revision {
				continue
			}
			if filter.Workspace != "" && filter.Workspace != pkgRev.Spec.PackageID.Workspace {
				continue
			}
			if filter.Lifecycle != "" && filter.Lifecycle != string(pkgRev.Spec.Lifecycle) {
				continue
			}
		}

		repo, err := r.getRepository(ctx, types.NamespacedName{Name: pkgid.GetRepoNameFromPkgRevName(pkgRev.Name), Namespace: pkgRev.Namespace})
		if err != nil {
			return nil, err
		}
		cachedRepo, err := r.cache.Open(ctx, repo)
		if err != nil {
			return nil, apierrors.NewInternalError(err)
		}
		resources, err := cachedRepo.GetResources(ctx, &pkgRev)
		if err != nil {
			return nil, apierrors.NewInternalError(err)
		}
		appendItem(v, buildPackageRevisionResources(&pkgRev, resources))
	}
	sort.SliceStable(newListObj.Items, func(i, j int) bool {
		return newListObj.Items[i].Name < newListObj.Items[j].Name
	})
	return newListObj, nil
}

func getListPrt(listObj runtime.Object) (reflect.Value, error) {
	listPtr, err := meta.GetItemsPtr(listObj)
	if err != nil {
		return reflect.Value{}, err
	}
	v, err := conversion.EnforcePtr(listPtr)
	if err != nil || v.Kind() != reflect.Slice {
		return reflect.Value{}, fmt.Errorf("need ptr to slice: %v", err)
	}
	return v, nil
}

func appendItem(v reflect.Value, obj runtime.Object) {
	v.Set(reflect.Append(v, reflect.ValueOf(obj).Elem()))
}
