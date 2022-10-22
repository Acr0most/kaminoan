package service

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var (
	appDir string
)

func init() {
	home, err := os.UserHomeDir()
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
			fmt.Printf(`Missing config file.
To create the file we need to know your default workspace directory.
If you want to clone repository absolute to the directory you're calling the command use empty string.
> `)
			if err := createDirAndConfigFile(); err != nil {
				panic(err)
				return err
			}

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			text = strings.Trim(text, " \n")

			viper.Set("Workspace", text)
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
