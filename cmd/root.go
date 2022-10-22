package cmd

import (
	"fmt"
	"github.com/Acr0most/kaminoan/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

const AppVersion = "0.0.1"

var (
	// Used for flags.
	cfgFile string
	verbose bool

	rootCmd = &cobra.Command{
		Use:   "kaminoan",
		Short: "Easy to use cli tool for organize your repositories.",
		Long:  "cli to wrap a normal `git clone <repository-url>` command and add some more useful logic",
		Args:  cobra.ArbitraryArgs,
		Run:   root,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kaminoan/settings.json)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "", false, "verbose")
	rootCmd.Flags().BoolP("version", "v", false, "Show the version and exit.")

	cobra.OnInitialize(initConfig)
}

func root(cmd *cobra.Command, args []string) {
	version, err := cmd.Flags().GetBool("version")

	if err == nil && version != false {
		fmt.Println("kaminoan " + AppVersion)
		os.Exit(0)
	}

	if verbose {
		viper.Set("verbose", true)
	}

	if len(args) > 0 {
		clone(cmd, args)
		os.Exit(0)
	}

	cmd.Help()
}

func initConfig() {
	cobra.CheckErr(service.InitConfig(cfgFile))
}
