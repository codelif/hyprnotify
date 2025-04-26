package internal

import (
	"fmt"
	"path"
	// "log"
	"github.com/spf13/viper"
)

type action struct {
	Silence bool
	Ignore bool
}

type RuleEntry struct {
	Matches string
	Action action
}

type config struct {
	Rule []RuleEntry
}

var (
	Config config
)

func (c *config) Load(location string) {
	if location != "" {
		fileDir := path.Dir(location)
		fileName := path.Base(location)

		viper.SetConfigName(fileName)
		viper.AddConfigPath(fileDir)
	} else {
		viper.SetConfigName("hyprnotify")
		viper.AddConfigPath("$HOME/.config/hypr")
	}
	
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Printf("Config file not found: %s \n", err)
		} else {
			fmt.Printf("Fatal error in config file: %s \n", err)
		}
		*c = config{}
	}
	
	if err := viper.Unmarshal(c); err != nil {
		fmt.Printf("Could not unmarshal the config file: %s \n", err)
	}
}
