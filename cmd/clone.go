package cmd

import (
	"github.com/Acr0most/kaminoan/model"
	"github.com/Acr0most/kaminoan/service"
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "clones a repository",
	Long:  "clones a repository",
	Args:  cobra.MinimumNArgs(1),
	Run:   clone,
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}

func clone(cmd *cobra.Command, args []string) {
	url := model.NewUrl(args[0])

	cloner := service.Kaminoan{}
	cloner.Clone(url)
}
