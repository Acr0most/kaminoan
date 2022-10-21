package main

import (
	"github.com/Acr0most/kaminoan/model"
	"github.com/Acr0most/kaminoan/service"
	"log"
	"os"
)

var (
	config *model.Config
)

func init() {
	config = model.NewConfig()
}

func main() {
	if !config.Load() {
		log.Println("missing config file.")
		return
	}

	if len(os.Args) < 2 {
		log.Println("missing repository url.")
		return
	}

	url := model.NewUrl(os.Args[1])

	if !url.Valid() {
		log.Println("invalid url format.")
		return
	}

	if len(os.Args) >= 3 {
		config.WorkspaceOverride = os.Args[2]
	}

	cloner := service.Kaminoan{}
	cloner.Clone(url, config)

}
