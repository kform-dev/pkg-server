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

package pkgserver

/*

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/henderiw/logger/log"
	"github.com/henderiw/store"
	configv1alpha1 "github.com/kform-dev/pkg-server/apis/config/v1alpha1"
	pkgv1alpha1 "github.com/kform-dev/pkg-server/apis/pkg/v1alpha1"
	"github.com/kform-dev/pkg-server/pkg/cache"
	watchermanager "github.com/kform-dev/pkg-server/pkg/watcher-manager"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type packageCommon struct {
	client         client.Client
	scheme         *runtime.Scheme
	gr             schema.GroupResource
	cache          *cache.Cache
	isNamespaced   bool
	newFunc        func() runtime.Object
	newListFunc    func() runtime.Object
	watcherManager watchermanager.WatcherManager
}

func (r *packageCommon) get(
	ctx context.Context,
	name string,
	options *metav1.GetOptions,
) (runtime.Object, error) {
	// get the repository from the name
	repo, err := r.getRepositoryFromName(ctx, name)
	if err != nil {
		return nil, err
	}
	cachedRepo, err := r.cache.Open(ctx, repo)
	if err != nil {
		return nil, apierrors.NewInternalError(err)
	}
	pkgRev, err := cachedRepo.GetPackageRevision(ctx, types.NamespacedName{Namespace: repo.Namespace, Name: name})
	if err != nil {
		return nil, apierrors.NewNotFound(r.gr, name)
	}
	return pkgRev, nil
}

type lister interface {
	list(ctx context.Context, options *metainternalversion.ListOptions) (runtime.Object, error)
}

func (r *packageCommon) list(ctx context.Context, options *metainternalversion.ListOptions) (runtime.Object, error) {
	//log := log.FromContext(ctx)

	packageFilter, err := parsePackageFieldSelector(options.FieldSelector)
	if err != nil {
		return nil, err
	}

	repoList, err := r.listRepositories(ctx, packageFilter)
	if err != nil {
		return nil, err
	}

	newListObj := r.newListFunc()
	v, err := getListPrt(newListObj)
	if err != nil {
		return nil, err
	}

	for _, repo := range repoList.Items {
		cachedRepo, err := r.cache.Open(ctx, &repo)
		if err != nil {
			return nil, apierrors.NewInternalError(err)
		}
		cachedRepo.ListPackageRevisions(ctx, func(ctx context.Context, key store.Key, obj *pkgv1alpha1.PackageRevision) {
			if packageFilter != nil {
				filter := true
				if packageFilter.Name != "" {
					if obj.Name == packageFilter.Name {
						filter = false
					} else {
						filter = true
					}
				}
				if packageFilter.Package != "" {
					if obj.Spec.Package == packageFilter.Package {
						filter = false
					} else {
						filter = true
					}
				}
				if packageFilter.Revision != "" {
					if obj.Spec.Revision == packageFilter.Revision {
						filter = false
					} else {
						filter = true
					}
				}
				if packageFilter.Workspace != "" {
					if obj.Spec.Workspace == packageFilter.Workspace {
						filter = false
					} else {
						filter = true
					}
				}
				if !filter {
					appendItem(v, obj)
				}

			} else {
				appendItem(v, obj)
			}
		})
	}
	return newListObj, nil
}

func (r *packageCommon) watch(
	ctx context.Context,
	options *metainternalversion.ListOptions,
) *watcher {

	// logger
	log := log.FromContext(ctx)

	if options.FieldSelector == nil {
		log.Info("watch", "options", *options, "fieldselector", "nil")
	} else {
		requirements := options.FieldSelector.Requirements()
		log.Info("watch", "options", *options, "fieldselector", options.FieldSelector.Requirements())
		for _, requirement := range requirements {
			log.Info("watch requirement",
				"Operator", requirement.Operator,
				"Value", requirement.Value,
				"Field", requirement.Field,
			)
		}
	}

	ctx, cancel := context.WithCancel(ctx)

	w := &watcher{
		cancel:         cancel,
		resultChan:     make(chan watch.Event),
		watcherManager: r.watcherManager,
	}

	go w.listAndWatch(ctx, r, options)

	return w
}

func (r *packageCommon) getRepositoryFromName(ctx context.Context, name string) (*configv1alpha1.Repository, error) {
	ns, namespaced := genericapirequest.NamespaceFrom(ctx)
	if namespaced != r.isNamespaced {
		return nil, apierrors.NewBadRequest(fmt.Errorf("namespace mismatch got %t, want %t", namespaced, r.isNamespaced).Error())
	}

	repositoryName, err := parseRepositoryName(name)
	if err != nil {
		return nil, apierrors.NewNotFound(r.gr, name)
	}

	return r.getRepository(ctx, types.NamespacedName{Name: repositoryName, Namespace: ns})
}

func (r *packageCommon) getRepository(ctx context.Context, nsn types.NamespacedName) (*configv1alpha1.Repository, error) {
	repo := configv1alpha1.Repository{}
	if err := r.client.Get(ctx, nsn, &repo); err != nil {
		return nil, apierrors.NewInternalError(err)
	}
	return &repo, nil
}

func (r *packageCommon) listRepositories(ctx context.Context, filter *packageFilter) (*configv1alpha1.RepositoryList, error) {
	opts := []client.ListOption{}
	if filter != nil {
		if filter.Namespace != "" {
			opts = append(opts, client.InNamespace(filter.Namespace))
		}
		if filter.Repository != "" {
			opts = append(opts, client.MatchingFields(map[string]string{
				"spec.repository": filter.Repository}))
		}
	}

	repoList := configv1alpha1.RepositoryList{}
	if err := r.client.List(ctx, &repoList, opts...); err != nil {
		return nil, apierrors.NewInternalError(err)
	}
	return &repoList, nil
}

func parseRepositoryName(name string) (string, error) {
	parts := strings.Split(name, ":")
	if len(parts) != 2 {
		return "", fmt.Errorf("malformed package revision name; expected <repo>:<package>:<revision> %q", name)
	}
	return parts[0], nil
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
*/