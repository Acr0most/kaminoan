package service

import (
	"errors"
	"fmt"
	"github.com/Acr0most/kaminoan/model"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Kaminoan struct{}

func (t *Kaminoan) Clone(repository *model.Repository, config *model.Config) {
	path := config.Workspace
	if !strings.HasSuffix(config.Workspace, "/") {
		path += "/"
	}
	path += repository.Path()

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		if mkdirErr := os.MkdirAll(path, os.ModePerm); mkdirErr != nil {
			log.Println("aborted.", mkdirErr)
			return
		}
	}

	cmd := exec.Command("git", "clone", repository.Url(), path)
	_, err := cmd.Output()

	if err != nil {
		fmt.Println("aborting.")
		return
	}

	log.Println("use: cd " + path)
}
