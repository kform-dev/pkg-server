package createcmd

import (
	"context"
	"fmt"

	//docs "github.com/kform-dev/pkg-server/internal/docs/generated/initdocs"

	"github.com/fatih/color"
	"github.com/henderiw/logger/log"
	"github.com/kform-dev/pkg-server/cmd/pkgctl/apis"
	"github.com/kform-dev/pkg-server/pkg/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// NewRunner returns a command runner.
func NewRunner(ctx context.Context, version string, cfg *genericclioptions.ConfigFlags, k8s bool) *Runner {
	r := &Runner{}
	cmd := &cobra.Command{
		Use:  "create NAME URL [flags]",
		Args: cobra.ExactArgs(2),
		//Short:   docs.InitShort,
		//Long:    docs.InitShort + "\n" + docs.InitLong,
		//Example: docs.InitExamples,
		PreRunE: r.preRunE,
		RunE:    r.runE,
	}

	r.Command = cmd
	r.cfg = cfg
	r.k8s = k8s

	r.Command.Flags().StringVarP(
		&r.secret, "secret", "", "", "secret used for accessing the repository")
	r.Command.Flags().BoolVarP(
		&r.deployment, "deployment", "d", false, "tags the repository as a deployment repository. packages in a deployment repository are considered for deployment dependeing on their lifecycle status")
	r.Command.Flags().StringVarP(
		&r.directory, "directory", "", "", "the directory withing the repository")
	return r
}

func NewCommand(ctx context.Context, version string, kubeflags *genericclioptions.ConfigFlags, k8s bool) *cobra.Command {
	return NewRunner(ctx, version, kubeflags, k8s).Command
}

type Runner struct {
	Command *cobra.Command
	cfg     *genericclioptions.ConfigFlags
	client  client.Client
	k8s   bool
	// dynamic input
	secret     string
	deployment bool
	directory  string
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
	log.Info("create repository", "k8s", r.k8s)
	if !r.k8s {
		repoName := args[0]
		repo := apis.Repo{
			URL:        args[1],
			Secret:     r.secret,
			Deployment: r.deployment,
			Directory:  r.directory,
		}
		log.Info("create repository", "repoName", repoName)
		viper.Set(fmt.Sprintf("repos.%s", repoName), repo)
		if err := viper.WriteConfig(); err != nil {
			color.Red("Error writing config file: %s", err.Error())
			return err
		}
	}

	return nil
}
