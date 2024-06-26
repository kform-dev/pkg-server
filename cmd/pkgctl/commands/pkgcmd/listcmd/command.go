package listcmd

import (
	"context"
	"fmt"
	"sort"

	//docs "github.com/kform-dev/pkg-server/internal/docs/generated/initdocs"

	pkgv1alpha1 "github.com/kform-dev/pkg-server/apis/pkg/v1alpha1"
	"github.com/kform-dev/pkg-server/pkg/client"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// NewRunner returns a command runner.
func NewRunner(ctx context.Context, version string, cfg *genericclioptions.ConfigFlags, k8s bool) *Runner {
	r := &Runner{}
	cmd := &cobra.Command{
		Use:  "list [flags]",
		Args: cobra.ExactArgs(0),
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

func NewCommand(ctx context.Context, version string, kubeflags *genericclioptions.ConfigFlags, k8s bool) *cobra.Command {
	return NewRunner(ctx, version, kubeflags, k8s).Command
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
	//log.Info("get packagerevision", "name", args[0])

	pkgRevList := &pkgv1alpha1.PackageRevisionList{}
	if err := r.client.List(ctx, pkgRevList); err != nil {
		return err
	}
	pkgRevs := make([]pkgv1alpha1.PackageRevision, len(pkgRevList.Items))
	copy(pkgRevs, pkgRevList.Items)
	sort.SliceStable(pkgRevs, func(i, j int) bool {
		return pkgRevs[i].Name < pkgRevs[j].Name
	})

	for _, pkgRev := range pkgRevs {
		fmt.Println(pkgRev.Name, pkgRev.Spec.PackageID.Package, pkgRev.Spec.PackageID.Workspace, pkgRev.Spec.PackageID.Revision, pkgRev.Spec.Lifecycle)
	}

	return nil
}
