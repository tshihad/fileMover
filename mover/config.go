package mover

import (
	"log"

	"github.com/spf13/viper"
)

func readConfig() (*config, error) {
	log.Println("Reading conf.toml")
	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	c := &config{}

	c.Dest = viper.GetString(destination)
	c.Source = viper.GetString(source)
	c.File = viper.GetString(filename)
	return c, nil
}
