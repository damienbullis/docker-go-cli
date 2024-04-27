package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Key1   string `mapstructure:"key1"`
	Key2   int    `mapstructure:"key2"`
	Nested struct {
		Key3 bool `mapstructure:"key3"`
	} `mapstructure:"nested"`
}

func main() {
	// set defaults
	viper.SetDefault("key1", "default1")
	viper.SetDefault("key2", 100)
	viper.SetDefault("nested.key3", true)

	// set config file
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	// read config file
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
	}

	// initialize config struct
	var config Config

	// unmarshal config into struct
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Error unmarshalling config: %s\n", err)
	}

	// print config
	fmt.Printf("Key1: %s\n", config.Key1)
	fmt.Printf("Key2: %d\n", config.Key2)
	fmt.Printf("Key3: %t\n", config.Nested.Key3)
	fmt.Printf("Key1: %s\n", config.Key1)
	fmt.Printf("Key2: %d\n", config.Key2)
	fmt.Printf("Key3: %t\n", config.Nested.Key3)

}
