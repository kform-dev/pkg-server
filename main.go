package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/henderiw/apiserver-builder/pkg/builder"
	"github.com/henderiw/logger/log"
	configv1alpha1 "github.com/kform-dev/pkg-server/apis/config/v1alpha1"
	"github.com/kform-dev/pkg-server/apis/generated/clientset/versioned"
	"github.com/kform-dev/pkg-server/apis/generated/clientset/versioned/scheme"
	pkgopenapi "github.com/kform-dev/pkg-server/apis/generated/openapi"
	pkgv1alpha1 "github.com/kform-dev/pkg-server/apis/pkg/v1alpha1"
	"github.com/kform-dev/pkg-server/pkg/auth/secret"
	"github.com/kform-dev/pkg-server/pkg/auth/ui"
	"github.com/kform-dev/pkg-server/pkg/cache"
	"github.com/kform-dev/pkg-server/pkg/pkgserver/packagerevision"
	"github.com/kform-dev/pkg-server/pkg/pkgserver/packagerevisionresource"
	"github.com/kform-dev/pkg-server/pkg/reconcilers"
	_ "github.com/kform-dev/pkg-server/pkg/reconcilers/all"
	"github.com/kform-dev/pkg-server/pkg/reconcilers/ctrlconfig"
	"github.com/kform-dev/pkg-server/pkg/reconcilers/packagediscovery/catalog"
	"go.uber.org/zap/zapcore"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/component-base/logs"
	configsyncv1beta1 "kpt.dev/configsync/pkg/api/configsync/v1beta1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

const (
	cacheDir              = "/cache"
	defaultEtcdPathPrefix = "/registry/pkg.kform.dev"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	l := log.NewLogger(&log.HandlerOptions{Name: "pkg-server-logger", AddSource: false})
	slog.SetDefault(l)
	ctx := log.IntoContext(context.Background(), l)
	log := log.FromContext(ctx)

	opts := zap.Options{
		TimeEncoder: zapcore.RFC3339NanoTimeEncoder,
	}

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	// setup controllers
	runScheme := runtime.NewScheme()
	if err := scheme.AddToScheme(runScheme); err != nil {
		log.Error("cannot initialize schema", "error", err)
		os.Exit(1)
	}
	// add the core object to the scheme
	for _, api := range (runtime.SchemeBuilder{
		clientgoscheme.AddToScheme,
		configv1alpha1.AddToScheme,
		pkgv1alpha1.AddToScheme,
		configsyncv1beta1.AddToScheme,
	}) {
		if err := api(runScheme); err != nil {
			log.Error("cannot add scheme", "err", err)
			os.Exit(1)
		}
	}
	runScheme.AddFieldLabelConversionFunc(
		schema.GroupVersionKind{
			Group:   pkgv1alpha1.Group,
			Version: pkgv1alpha1.Version,
			Kind:    pkgv1alpha1.PackageRevisionKind,
		},
		pkgv1alpha1.ConvertPackageRevisionsFieldSelector,
	)
	runScheme.AddFieldLabelConversionFunc(
		schema.GroupVersionKind{
			Group:   pkgv1alpha1.Group,
			Version: pkgv1alpha1.Version,
			Kind:    pkgv1alpha1.PackageRevisionResourcesKind,
		},
		pkgv1alpha1.ConvertPackageRevisionResourcesFieldSelector,
	)

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: runScheme,
	})
	if err != nil {
		log.Error("cannot start manager", "err", err)
		os.Exit(1)
	}

	clientset, err := versioned.NewForConfig(mgr.GetConfig())
	if err != nil {
		panic(err.Error())
	}

	cache := cache.NewCache(cacheDir, 1*time.Minute, cache.Options{
		Client: mgr.GetClient(),
		CredentialResolver: secret.NewCredentialResolver(mgr.GetClient(), []secret.Resolver{
			secret.NewBasicAuthResolver(),
		}),
		UserInfoProvider: &ui.ApiserverUserInfoProvider{},
	})

	ctrlCfg := &ctrlconfig.ControllerConfig{
		RepoCache:    cache,
		CatalogStore: catalog.NewStore(),
		ClientSet:    clientset,
	}
	for name, reconciler := range reconcilers.Reconcilers {
		log.Info("reconciler", "name", name, "enabled", IsReconcilerEnabled(name))
		if IsReconcilerEnabled(name) {
			_, err := reconciler.SetupWithManager(ctx, mgr, ctrlCfg)
			if err != nil {
				log.Error("cannot add controllers to manager", "err", err.Error())
				os.Exit(1)
			}
		}
	}

	go func() {
		if err := builder.APIServer.
			WithServerName("pkg-server").
			WithEtcdPath(defaultEtcdPathPrefix).
			WithOpenAPIDefinitions("Package", "v0.0.0", pkgopenapi.GetOpenAPIDefinitions).
			WithResourceAndHandler(ctx, &pkgv1alpha1.PackageRevision{}, packagerevision.NewProvider(ctx, mgr.GetClient(), cache)).
			WithResourceAndHandler(ctx, &pkgv1alpha1.PackageRevisionResources{}, packagerevisionresource.NewProvider(ctx, mgr.GetClient(), cache)).
			Execute(ctx); err != nil {
			log.Info("cannot start pkg-server")
		}
	}()

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		log.Error("unable to set up health check", "error", err.Error())
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		log.Error("unable to set up ready check", "error", err.Error())
		os.Exit(1)
	}

	log.Info("starting manager")
	if err := mgr.Start(ctx); err != nil {
		log.Error("problem running manager", "error", err.Error())
		os.Exit(1)
	}
}

func IsReconcilerEnabled(reconcilerName string) bool {
	if _, found := os.LookupEnv(fmt.Sprintf("ENABLE_%s", strings.ToUpper(reconcilerName))); found {
		return true
	}
	return false
}
