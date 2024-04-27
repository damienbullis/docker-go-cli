package core

import "fmt"

type Config struct {
	Key1   string `mapstructure:"key1"`
	Key2   int    `mapstructure:"key2"`
	Nested struct {
		Key3 bool `mapstructure:"key3"`
	} `mapstructure:"nested"`
}

func PrintConfig(config *Config) {

	fmt.Printf("Key1: %s\n", config.Key1)
	fmt.Printf("Key2: %d\n", config.Key2)
	fmt.Printf("Key3: %t\n", config.Nested.Key3)

}
