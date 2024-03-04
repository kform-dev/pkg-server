package deletecmd

import (
	"context"

	//docs "github.com/kform-dev/pkg-server/internal/docs/generated/initdocs"

	pkgv1alpha1 "github.com/kform-dev/pkg-server/apis/pkg/v1alpha1"
	"github.com/kform-dev/pkg-server/pkg/client"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// NewRunner returns a command runner.
func NewRunner(ctx context.Context, version string, cfg *genericclioptions.ConfigFlags) *Runner {
	r := &Runner{}
	cmd := &cobra.Command{
		Use:  "delete PKGREV [flags]",
		Args: cobra.ExactArgs(1),
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
	//log.Info("delete packagerevision", "name", args[0])

	pkgRev := &pkgv1alpha1.PackageRevision{}
	if err := r.client.Get(ctx, types.NamespacedName{Namespace: "default", Name: args[0]}, pkgRev); err != nil {
		return err
	}
	pkgRev.Spec.Tasks = []pkgv1alpha1.Task{}
	return r.client.Delete(ctx, pkgRev)
}
