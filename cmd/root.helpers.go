package cmd

import (
	"go-commander/utils"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

func loadConfigFile() {
	// if config file path is provided, set it
	if configFilePath != "" {
		log.Debugf("Checking config: %s", configFilePath)
		viper.SetConfigFile(configFilePath)

		// read config file
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		}

		log.Infof("%s %s found!", utils.Icon.Mag, viper.ConfigFileUsed())
	} else {

		// set defaults config file
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		log.Debug("Checking for default config file")

		// read config file
		err := viper.ReadInConfig()
		if err != nil {
			switch err.(type) {
			case viper.ConfigFileNotFoundError:
				log.Debug("Config file not found, using default values")
			default:
				log.Fatal(err)
			}
		}
	}
}

func setConfigsDefaults() {
	log.Debug("Setting defaults")
	viper.SetDefault("key1", "default1")
	viper.SetDefault("key2", 100)
	viper.SetDefault("nested.key3", true)
}
