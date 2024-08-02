package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

var gViper *viper.Viper

func Init(configName, configType, configPath string) error {
	gViper = viper.GetViper()
	// gViper.SetConfigFile() // name of config file (without extension)
	gViper.SetConfigType(configType) // REQUIRED if the config file does not have the extension in the name
	gViper.SetConfigName(configName) // name of config file (without extension)
	gViper.AddConfigPath(configPath) // optionally look for config in the working directory

	// Handle errors reading the config file
	err := gViper.ReadInConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return fmt.Errorf("配置文件未找到: %v", err)
		}
	}

	return nil
}

func InitForFile(configFile string) error {
	gViper = viper.GetViper()
	gViper.SetConfigFile(configFile) // name of config file (without extension)

	// Handle errors reading the config file
	err := gViper.ReadInConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return fmt.Errorf("配置文件未找到: %v", err)
		}
	}

	return nil
}

func GetViper() *viper.Viper {
	return gViper
}
