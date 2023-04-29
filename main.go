package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func main() {
	// Find the current user

	// call multiple times to add many search paths
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")

	viper.SetConfigName(".realtalk")
	viper.SetConfigType("yaml")

	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})

	// Read in the config file
	if err := viper.ReadInConfig(); err != nil {
		// If the config file doesn't exist, create it with default values
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
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
		} else {
			fmt.Println("Error reading config file:", err)
			os.Exit(1)
		}
	}

	// Get values from the config file
	ContentDir := viper.GetString("ContentDir")
	LayoutDir := viper.GetString("LayoutDir")

	// Print out the values
	fmt.Println("ContentDir:", ContentDir)
	fmt.Println("LayoutDir:", LayoutDir)
}
