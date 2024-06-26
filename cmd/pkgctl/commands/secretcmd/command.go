package secretcmd

import (
	"context"

	//docs "github.com/kform-dev/pkg-server/internal/docs/generated/initdocs"

	"github.com/kform-dev/pkg-server/cmd/pkgctl/commands/secretcmd/createcmd"
	"github.com/kform-dev/pkg-server/cmd/pkgctl/commands/secretcmd/deletecmd"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// NewRunner returns a command runner.
func GetCommand(ctx context.Context, version string, cfg *genericclioptions.ConfigFlags, k8s bool) *cobra.Command {
	cmd := &cobra.Command{
		Use: "secret",
		//Short:   docs.InitShort,
		//Long:    docs.InitShort + "\n" + docs.InitLong,
		//Example: docs.InitExamples,
		RunE: func(cmd *cobra.Command, args []string) error {
			h, err := cmd.Flags().GetBool("help")
			if err != nil {
				return err
			}
			if h {
				return cmd.Help()
			}
			return cmd.Usage()
		},
	}

	cmd.AddCommand(
		createcmd.NewCommand(ctx, version, cfg, k8s),
		deletecmd.NewCommand(ctx, version, cfg, k8s),
	)
	return cmd
}
