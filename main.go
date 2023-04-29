package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func main() {
	setupConfig()

	printConfig("ContentDir", "LayoutDir", "Taxonomies.tag")
}

func setupConfig() {
	addConfigPaths()
	setConfigNameAndType()
	setDefaultValues()

	if err := viper.ReadInConfig(); err != nil {
		handleConfigReadError(err)
	}
}

func addConfigPaths() {
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")
}

func setConfigNameAndType() {
	viper.SetConfigName(".realtalk")
	viper.SetConfigType("yaml")
}

func setDefaultValues() {
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
}

func handleConfigReadError(err error) {
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		createConfigFile()
	} else {
		fmt.Println("Error reading config file:", err)
		os.Exit(1)
	}
}

func createConfigFile() {
	fmt.Println("Config file not found, creating it with default values...")

	if err := viper.SafeWriteConfig(); err != nil {
		if os.IsNotExist(err) {
			if err := viper.WriteConfig(); err != nil {
				fmt.Println("Error writing config file:", err)
				os.Exit(1)
			}
		}
	}

	fmt.Println("Config file created with default values.")
}

func printConfig(keys ...string) {
	for _, key := range keys {
		value := viper.Get(key)
		fmt.Printf("%s: %v\n", key, value)
	}
}
