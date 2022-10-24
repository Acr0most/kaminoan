package service

import (
	"errors"
	"fmt"
	prompt "github.com/Acr0most/kaminoan/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	home   string
	appDir string
)

func init() {
	var err error
	home, err = os.UserHomeDir()
	cobra.CheckErr(err)
	appDir = home + "/.kaminoan/"
}

func InitConfig(cfgFile string) error {

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory / .kaminoan with name "settings" (without extension).
		viper.AddConfigPath(appDir)
		viper.SetConfigType("json")
		viper.SetConfigName("settings")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			workspace := prompt.Prompt(`Missing config file.
To create the file we need to know your default workspace directory.
If you want to clone repository absolute to the directory you're calling the command use empty string.`, "")

			if err := createDirAndConfigFile(); err != nil {
				return err
			}

			viper.Set("workspace", workspace)
			viper.Set("auth.private_key_file", home+"/.ssh/id_rsa")
			viper.Set("auth.private_key_requires_password", true)

			return viper.WriteConfig()
		} else {
			return err
		}
	}

	return nil
}

func createDirAndConfigFile() error {
	if _, err := os.Stat(appDir); errors.Is(err, os.ErrNotExist) {
		_ = os.MkdirAll(appDir, os.ModePerm)
	}

	return os.WriteFile(appDir+"/"+"settings.json", []byte{}, os.ModePerm)
}
