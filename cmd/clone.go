package cmd

import (
	"github.com/Acr0most/kaminoan/model"
	"github.com/Acr0most/kaminoan/service"
	"github.com/spf13/cobra"
	"log"
	"os"
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

	if !url.Valid() {
		log.Println("Invalid repository url")
		os.Exit(1)
	}

	cloner := service.Kaminoan{}
	cobra.CheckErr(cloner.Clone(url))
}
