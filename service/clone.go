package service

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/Acr0most/kaminoan/model"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"strings"
)

type Kaminoan struct{}

func (t *Kaminoan) Clone(repository *model.Repository) error {
	path := viper.GetString("workspace")
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	path += repository.Path()

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		if mkdirErr := os.MkdirAll(path, os.ModePerm); mkdirErr != nil {
			log.Println("aborted.", mkdirErr)
			return mkdirErr
		}
	}

	var progress io.Writer
	if viper.GetBool("verbose") {
		progress = os.Stdout
	}

	var auth transport.AuthMethod
	if repository.Mode() == model.SSH {
		var err error
		auth, err = getSSHAuth()
		if err != nil {
			return err
		}
	}

	_, err := git.PlainClone(path, false, &git.CloneOptions{
		Auth:     auth,
		URL:      repository.Url(),
		Progress: progress,
	})

	if err != nil {
		if errors.Is(err, git.ErrRepositoryAlreadyExists) {
			return perhapsUpdateRepository(path, repository)
		}

		return err
	}

	log.Println("use: cd " + path)
	return nil
}

func perhapsUpdateRepository(path string, repository *model.Repository) error {
	fmt.Printf("Repository %s already exists. Update [y/n]?\n", repository.Url())
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, " \n")

	if text == "y" {
		repo, err := git.PlainOpen(path)
		if err != nil {
			return err
		}

		worktree, err := repo.Worktree()
		if err != nil {
			return err
		}

		err = worktree.Pull(&git.PullOptions{})
		if errors.Is(err, git.NoErrAlreadyUpToDate) {
			log.Println("Already up to date")
			os.Exit(0)
		}

		return err
	}

	return nil
}

func getSSHAuth() (transport.AuthMethod, error) {
	password := viper.GetString("auth.private_key_password")
	privateKeyFile := viper.GetString("auth.private_key_file")

	_, err := os.Stat(privateKeyFile)
	if err != nil {
		return nil, err
	}

	publicKeys, err := ssh.NewPublicKeysFromFile("git", privateKeyFile, password)
	if err != nil {
		return nil, err
	}

	return publicKeys, nil
}
