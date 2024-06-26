package getcmd

import (
	"context"
	"fmt"

	//docs "github.com/kform-dev/pkg-server/internal/docs/generated/initdocs"

	pkgv1alpha1 "github.com/kform-dev/pkg-server/apis/pkg/v1alpha1"
	"github.com/kform-dev/pkg-server/pkg/client"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"sigs.k8s.io/yaml"
)

// NewRunner returns a command runner.
func NewRunner(ctx context.Context, version string, cfg *genericclioptions.ConfigFlags, k8s bool) *Runner {
	r := &Runner{}
	cmd := &cobra.Command{
		Use:  "get PKGREV [flags]",
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

	namespace := "default"
	if r.cfg.Namespace != nil && *r.cfg.Namespace != "" {
		namespace = *r.cfg.Namespace
	}

	pkgRev := &pkgv1alpha1.PackageRevision{}
	if err := r.client.Get(ctx, types.NamespacedName{Namespace: namespace, Name: args[0]}, pkgRev); err != nil {
		return err
	}
	b, err := yaml.Marshal(pkgRev)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}
