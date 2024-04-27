package cmd

import (
	"go-commander/core"
	"go-commander/utils"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "Go Commander",
	Short: "Go Commander is a CLI application testing cobra / viper / docker",
	Run:   run,
}

var configFilePath string
var verbose bool

func init() {
	rootCmd.Flags().StringVarP(&configFilePath, "config", "c", "", "config file path")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func run(cmd *cobra.Command, args []string) {
	if verbose {
		log.SetLevel(log.DebugLevel)
		log.Debug("Verbose output enabled")
	}

	// if config file path is provided, set it
	log.Debugf("Config file provided: %s", configFilePath)
	if configFilePath != "" {
		viper.SetConfigFile(configFilePath)

		// read config file
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		}

		log.Infof("🔍 %s found!", viper.ConfigFileUsed())
	} else {

		log.Debugf("Using default config")
		// set defaults config file
		viper.SetConfigName("config")
		viper.AddConfigPath(".")

		// read config file
		err := viper.ReadInConfig()
		if err != nil {
			switch err.(type) {
			case viper.ConfigFileNotFoundError:
				log.Info("Config file not found, using defaults")
			default:
				log.Fatal(err)
			}
		}
	}

	// set defaults
	log.Debug("Setting defaults")
	viper.SetDefault("key1", "default1")
	viper.SetDefault("key2", 100)
	viper.SetDefault("nested.key3", true)

	// initialize config struct
	var config core.Config

	// unmarshal config into struct
	log.Debug("Unmarshalling config")
	if err := viper.Unmarshal(&config); err != nil {
		log.Errorf("Error unmarshalling config: %s\n", err)
		log.Warn("Falling back to defaults")
	} else {
		log.Debugf("Config successfully unmarshalled! %s", utils.Icon.Tada)
	}

	// print config
	log.Debug("Starting")
	core.PrintConfig(&config)

	log.Infof("Done! %s", utils.Icon.Tada)
}
