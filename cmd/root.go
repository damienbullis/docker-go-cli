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
	utils.InitializeLog(verbose)

	loadConfigFile()

	setConfigsDefaults()

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

	log.Debug("Starting...")
	core.PrintConfig(&config)

	log.Infof("Done! %s", utils.Icon.Tada)
}
