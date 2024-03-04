package commands

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/kform-dev/pkg-server/cmd/pkgctl/commands/approvecmd"
	"github.com/kform-dev/pkg-server/cmd/pkgctl/commands/createcmd"
	"github.com/kform-dev/pkg-server/cmd/pkgctl/commands/deletecmd"
	"github.com/kform-dev/pkg-server/cmd/pkgctl/commands/draftcmd"
	"github.com/kform-dev/pkg-server/cmd/pkgctl/commands/getcmd"
	"github.com/kform-dev/pkg-server/cmd/pkgctl/commands/listcmd"
	"github.com/kform-dev/pkg-server/cmd/pkgctl/commands/proposecmd"
	"github.com/kform-dev/pkg-server/cmd/pkgctl/commands/proposedeletecmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/rest"
)

const (
	defaultConfigFileSubDir  = "pkgctl"
	defaultConfigFileName    = "pkgctl"
	defaultConfigFileNameExt = "yaml"
	defaultConfigEnvPrefix   = "PKGCTL"
	//defaultDBPath            = "package_db"
)

var (
	configFile string
)

func GetMain(ctx context.Context) *cobra.Command {
	//showVersion := false
	cmd := &cobra.Command{
		Use:          "pkgctl",
		Short:        "pkgctl is a pkg management tool",
		Long:         "pkgctl is a pkg management tool",
		SilenceUsage: true,
		// We handle all errors in main after return from cobra so we can
		// adjust the error message coming from libraries
		SilenceErrors: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// initialize viper
			// ensure the viper config directory exists
			//cobra.CheckErr(fsys.EnsureDir(ctx, xdg.ConfigHome, defaultConfigFileSubDir))
			// initialize viper settings
			initConfig()
			// create package store if it does not exist
			//_, err := store.New(cmd.Context(),
			//	filepath.Join(xdg.ConfigHome, defaultConfigFileSubDir, defaultDBPath))
			//return err
			return nil
		},
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

	pf := cmd.PersistentFlags()

	kubeflags := genericclioptions.NewConfigFlags(true)
	kubeflags.AddFlags(pf)

	kubeflags.WrapConfigFn = func(rc *rest.Config) *rest.Config {
		rc.UserAgent = fmt.Sprintf("pkg/%s", version)
		return rc
	}

	cmd.AddCommand(createcmd.NewCommand(ctx, version, kubeflags))
	cmd.AddCommand(draftcmd.NewCommand(ctx, version, kubeflags))
	cmd.AddCommand(proposecmd.NewCommand(ctx, version, kubeflags))
	cmd.AddCommand(approvecmd.NewCommand(ctx, version, kubeflags))
	cmd.AddCommand(proposedeletecmd.NewCommand(ctx, version, kubeflags))
	cmd.AddCommand(deletecmd.NewCommand(ctx, version, kubeflags))
	cmd.AddCommand(getcmd.NewCommand(ctx, version, kubeflags))
	cmd.AddCommand(listcmd.NewCommand(ctx, version, kubeflags))
	//cmd.PersistentFlags().StringVar(&configFile, "config", "c", fmt.Sprintf("Default config file (%s/%s/%s.%s)", xdg.ConfigHome, defaultConfigFileSubDir, defaultConfigFileName, defaultConfigFileNameExt))

	return cmd
}

type Runner struct {
	Command *cobra.Command
	//Ctx     context.Context
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {

		viper.AddConfigPath(filepath.Join(xdg.ConfigHome, defaultConfigFileName, defaultConfigFileNameExt))
		viper.SetConfigType(defaultConfigFileNameExt)
		viper.SetConfigName(defaultConfigFileSubDir)

		_ = viper.SafeWriteConfig()
	}

	//viper.Set("kubecontext", kubecontext)
	//viper.Set("kubeconfig", kubeconfig)

	viper.SetEnvPrefix(defaultConfigEnvPrefix)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		_ = 1
	}
}
