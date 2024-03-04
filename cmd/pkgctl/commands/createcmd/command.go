package createcmd

import (
	"context"
	"fmt"

	//docs "github.com/kform-dev/pkg-server/internal/docs/generated/initdocs"

	pkgv1alpha1 "github.com/kform-dev/pkg-server/apis/pkg/v1alpha1"
	"github.com/kform-dev/pkg-server/apis/pkgid"
	"github.com/kform-dev/pkg-server/pkg/client"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// NewRunner returns a command runner.
func NewRunner(ctx context.Context, version string, cfg *genericclioptions.ConfigFlags) *Runner {
	r := &Runner{}
	cmd := &cobra.Command{
		Use:  "create SRC[<REPO>:<PKG>:<REVISION>] [<REPO>:<PKG>:<WORKSPACE>] [flags]",
		Args: cobra.ExactArgs(2),
		//Short:   docs.InitShort,
		//Long:    docs.InitShort + "\n" + docs.InitLong,
		//Example: docs.InitExamples,
		PreRunE: r.preRunE,
		RunE:    r.runE,
	}

	r.Command = cmd
	r.cfg = cfg
	/*
		r.Command.Flags().StringVar(
			&r.FnConfigDir, "fn-config-dir", "", "dir where the function config files are located")
	*/
	return r
}

func NewCommand(ctx context.Context, version string, kubeflags *genericclioptions.ConfigFlags) *cobra.Command {
	return NewRunner(ctx, version, kubeflags).Command
}

type Runner struct {
	Command *cobra.Command
	cfg     *genericclioptions.ConfigFlags
	client  client.Client
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
	//log := log.FromContext(ctx)
	//log.Info("create packagerevision", "src", args[0], "dst", args[1])

	srcPkgID, err := pkgid.ParsePkgID(args[0])
	if err != nil {
		return err
	}
	dstPkgID, err := pkgid.ParsePkgID(args[1])
	if err != nil {
		return err
	}

	pkgRev := pkgv1alpha1.BuildPackageRevision(
		metav1.ObjectMeta{
			Name:      args[1],
			Namespace: "default",
		},
		pkgv1alpha1.PackageRevisionSpec{
			PackageID: *dstPkgID,
			Lifecycle: pkgv1alpha1.PackageRevisionLifecycleDraft,
			Upstream: &pkgid.Upstream{
				Repository: srcPkgID.Repository,
				Realm:      srcPkgID.Realm,
				Package:    srcPkgID.Package,
				Revision:   srcPkgID.Revision,
			},
			Tasks: []pkgv1alpha1.Task{
				{
					Type: pkgv1alpha1.TaskTypeClone,
				},
			},
		},
		pkgv1alpha1.PackageRevisionStatus{},
	)
	if err := r.client.Create(ctx, pkgRev); err != nil {
		return err
	}
	fmt.Println(pkgRev.Name)
	return nil
}
