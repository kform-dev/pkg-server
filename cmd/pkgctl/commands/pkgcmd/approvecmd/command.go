package approvecmd

import (
	"context"
	"fmt"
	"path/filepath"

	//docs "github.com/kform-dev/pkg-server/internal/docs/generated/initdocs"

	"github.com/henderiw/logger/log"
	configv1alpha1 "github.com/kform-dev/pkg-server/apis/config/v1alpha1"
	pkgv1alpha1 "github.com/kform-dev/pkg-server/apis/pkg/v1alpha1"
	"github.com/kform-dev/pkg-server/apis/pkgid"
	"github.com/kform-dev/pkg-server/cmd/pkgctl/apis"
	"github.com/kform-dev/pkg-server/pkg/auth/ui"
	viperauth "github.com/kform-dev/pkg-server/pkg/auth/viper"
	"github.com/kform-dev/pkg-server/pkg/client"
	"github.com/kform-dev/pkg-server/pkg/git/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// NewRunner returns a command runner.
func NewRunner(ctx context.Context, version string, cfg *genericclioptions.ConfigFlags, k8s bool) *Runner {
	r := &Runner{}
	cmd := &cobra.Command{
		Use:  "approve PKGREV[<Target>.<REPO>.<REALM>.<PACKAGE>.<WORKSPACE>] [flags]",
		Args: cobra.MinimumNArgs(1),
		//Short:   docs.InitShort,
		//Long:    docs.InitShort + "\n" + docs.InitLong,
		//Example: docs.InitExamples,
		PreRunE: r.preRunE,
		RunE:    r.runE,
	}

	r.Command = cmd
	r.cfg = cfg
	r.k8s = k8s
	r.Command.Flags().StringVar(
		&r.revision, "revision", "", "revision of the package to be cloned")

	return r
}

func NewCommand(ctx context.Context, version string, kubeflags *genericclioptions.ConfigFlags, k8s bool) *cobra.Command {
	return NewRunner(ctx, version, kubeflags, k8s).Command
}

type Runner struct {
	Command  *cobra.Command
	cfg      *genericclioptions.ConfigFlags
	client   client.Client
	k8s      bool
	revision string
}

func (r *Runner) preRunE(_ *cobra.Command, _ []string) error {
	client, err := client.CreateClientWithFlags(r.cfg)
	if err != nil {
		return err
	}
	r.client = client
	return nil
}

func (r *Runner) runE(c *cobra.Command, args []string) error {
	ctx := c.Context()
	log := log.FromContext(ctx)
	log.Debug("approve packagerevision", "name", args[0])

	namespace := "default"
	if r.cfg.Namespace != nil && *r.cfg.Namespace != "" {
		namespace = *r.cfg.Namespace
	}

	pkgRevName := args[0]
	if r.k8s {

		// fetch the package revision
		pkgRev := &pkgv1alpha1.PackageRevision{}
		if err := r.client.Get(ctx, types.NamespacedName{Namespace: namespace, Name: pkgRevName}, pkgRev); err != nil {
			return err
		}

		pkgRev.Spec.Tasks = []pkgv1alpha1.Task{}
		pkgRev.Spec.Lifecycle = pkgv1alpha1.PackageRevisionLifecyclePublished
		return r.client.Update(ctx, pkgRev)
	} else {
		if r.revision == "" {
			return fmt.Errorf("a revision is required for approving a package client side")
		}
		pkgID, err := pkgid.ParsePkgRev2PkgID(pkgRevName)
		if err != nil {
			return err
		}
		pkgID.Revision = r.revision

		repoName := pkgID.Repository
		var repo apis.Repo
		if err := viper.UnmarshalKey(fmt.Sprintf("repos.%s", repoName), &repo); err != nil {
			return err
		}

		dir := pkgID.Package
		if len(args) > 2 {
			dir = args[2]
		}

		pkgDir := filepath.Join(dir, pkg.LocalGitDirectory)

		repoCR := configv1alpha1.BuildRepository(
			metav1.ObjectMeta{
				Namespace: namespace,
				Name:      repoName,
			},
			configv1alpha1.RepositorySpec{
				Type: configv1alpha1.RepositoryTypeGit,
				Git: &configv1alpha1.GitRepository{
					URL:         repo.URL,
					Credentials: repo.Secret,
					Directory:   repo.Directory,
				},
				Deployment: repo.Deployment,
			},
			configv1alpha1.RepositoryStatus{},
		)
		cachedRepo, err := pkg.OpenRepository(ctx, pkgDir, repoCR, &pkg.Options{
			CredentialResolver: viperauth.NewCredentialResolver(),
			UserInfoProvider:   &ui.ApiserverUserInfoProvider{},
		})
		if err != nil {
			return err
		}

		pkgRev := pkgv1alpha1.BuildPackageRevision(
			metav1.ObjectMeta{
				Namespace: namespace,
				Name:      pkgRevName,
			},
			pkgv1alpha1.PackageRevisionSpec{
				PackageID: *pkgID,
				Lifecycle: pkgv1alpha1.PackageRevisionLifecyclePublished,
			},
			pkgv1alpha1.PackageRevisionStatus{},
		)
		if err := cachedRepo.EnsurePackageRevision(ctx, pkgRev); err != nil {
			return err
		}

		return nil
	}
}
