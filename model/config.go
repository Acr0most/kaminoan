package model

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Config struct {
	Workspace  string `json:"workspace"`
	path       string
	configFile string
}

func NewConfig() *Config {
	dir, _ := os.UserHomeDir()

	return &Config{
		path:       dir + "/.kaminoan/",
		configFile: dir + "/.kaminoan/settings.json",
	}
}

func (t *Config) Load() bool {
	if _, err := os.Stat(t.configFile); errors.Is(err, os.ErrNotExist) {
		log.Print(fmt.Sprintf("Missing config file %s.\n\nTo create the file we need to know your default workspace directory.\nIf you want to clone repository absolute to the directory you're calling the command use empty string.", t.configFile))
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, " \n")

		if _, err := os.Stat(text); errors.Is(err, os.ErrNotExist) {
			log.Println("this directory doesn't exists. Please create and retry.")
		}

		t.Workspace = text
		t.Create()
	}

	jsonFile, err := os.Open(t.configFile)

	if err != nil {
		return false
	}

	var config Config
	content, _ := io.ReadAll(jsonFile)

	_ = json.Unmarshal(content, &config)

	t.Workspace = config.Workspace
	return true
}

func (t *Config) Write() bool {
	if _, err := os.Stat(t.configFile); errors.Is(err, os.ErrNotExist) {
		return false
	}

	jsonFile, _ := os.Open(t.configFile)
	content, _ := json.Marshal(t)
	_, _ = jsonFile.Write(content)

	return true
}

func (t *Config) Create() bool {
	if _, err := os.Stat(t.path); errors.Is(err, os.ErrNotExist) {
		_ = os.MkdirAll(t.path, os.ModePerm)
	}

	jsonFile, _ := os.Create(t.configFile)
	content, _ := json.Marshal(t)
	_, _ = jsonFile.Write(content)

	return true
}
