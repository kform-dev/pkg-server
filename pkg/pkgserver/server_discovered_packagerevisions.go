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

	builderrest "github.com/henderiw/apiserver-builder/pkg/builder/rest"
	pkgv1alpha1 "github.com/kform-dev/pkg-server/apis/pkg/v1alpha1"
	"github.com/kform-dev/pkg-server/pkg/cache"
	watchermanager "github.com/kform-dev/pkg-server/pkg/watcher-manager"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var tracer = otel.Tracer("pkg-server")

func NewDiscoveredPkgRevisionProviderHandler(ctx context.Context, client client.Client, cache *cache.Cache) builderrest.ResourceHandlerProvider {
	return func(ctx context.Context, scheme *runtime.Scheme, getter generic.RESTOptionsGetter) (rest.Storage, error) {
		return NewDiscoveredPkgRevREST(ctx, client, scheme, cache)
	}
}

func NewDiscoveredPkgRevREST(ctx context.Context, client client.Client, scheme *runtime.Scheme, cache *cache.Cache) (rest.Storage, error) {
	if err := pkgv1alpha1.AddToScheme(scheme); err != nil {
		return nil, err
	}

	r := &discoveredPkgRevs{
		packageCommon: packageCommon{
			client:       client,
			cache:        cache,
			isNamespaced: (&pkgv1alpha1.PackageRevision{}).NamespaceScoped(),
			newFunc:      func() runtime.Object { return &pkgv1alpha1.PackageRevision{} },
			newListFunc:  func() runtime.Object { return &pkgv1alpha1.PackageRevisionList{} },
			watcherManager: watchermanager.New(32),
		},
		TableConvertor: NewDiscoveredPackageRevisionTableConvertor((&pkgv1alpha1.PackageRevision{}).GetGroupVersionResource().GroupResource()),

	}
	go r.packageCommon.watcherManager.Start(ctx)
	return r, nil
}

var _ rest.Lister = &discoveredPkgRevs{}
var _ rest.Getter = &discoveredPkgRevs{}
var _ rest.Scoper = &discoveredPkgRevs{}
var _ rest.Storage = &discoveredPkgRevs{}
var _ rest.TableConvertor = &discoveredPkgRevs{}
var _ rest.SingularNameProvider = &discoveredPkgRevs{}

type discoveredPkgRevs struct {
	packageCommon
	rest.TableConvertor
}

func (r *discoveredPkgRevs) Destroy() {}

func (r *discoveredPkgRevs) New() runtime.Object {
	return r.newFunc()
}

func (r *discoveredPkgRevs) NewList() runtime.Object {
	return r.newListFunc()
}

func (r *discoveredPkgRevs) NamespaceScoped() bool {
	return r.isNamespaced
}

func (r *discoveredPkgRevs) GetSingularName() string {
	return "discoveredpackagerevision"
}

func (r *discoveredPkgRevs) Get(
	ctx context.Context,
	name string,
	options *metav1.GetOptions,
) (runtime.Object, error) {

	// Start OTEL tracer
	ctx, span := tracer.Start(ctx, "discoveredpackagerevisions::Get", trace.WithAttributes())
	defer span.End()

	return r.get(ctx, name, options)
}

func (r *discoveredPkgRevs) List(
	ctx context.Context,
	options *metainternalversion.ListOptions,
) (runtime.Object, error) {

	// Start OTEL tracer
	ctx, span := tracer.Start(ctx, "discoveredpackagerevisions::List", trace.WithAttributes())
	defer span.End()

	return r.list(ctx, options)
}

func (r *discoveredPkgRevs) Watch(
	ctx context.Context,
	options *metainternalversion.ListOptions,
) (watch.Interface, error) {
	// Start OTEL tracer
	ctx, span := tracer.Start(ctx, "discoveredpackagerevisions::Watch", trace.WithAttributes())
	defer span.End()

	w := r.watch(ctx, options)

	return w, nil
}
*/
