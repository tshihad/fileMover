package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

const (
	source      = "source"
	destination = "destination"
	filename    = "filename"
)

type config struct {
	Source string
	Dest   string
	File   string
}

func main() {

	c, err := readConfig()
	if err != nil {
		log.Fatal("error while reading configuration file")
	}

	for {
		moveFile(c)
		time.Sleep(time.Second * 100)
	}

}

func readConfig() (*config, error) {
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

func moveFile(c *config) error {

	files, err := ioutil.ReadDir(c.Source)
	if err != nil {
		return errors.New("Source destination location doesnt exists " + err.Error())
	}
	if len(files) > 0 {
		if err := checkDestinationDir(c); err != nil {
			return errors.New("Destination location doest exists " + err.Error())
		}
		oldLoc := c.Source + files[0].Name()
		newLoc := c.Dest + c.File
		err := os.Rename(oldLoc, newLoc)
		if err != nil {
			return errors.New("Failed to move files " + err.Error())
		}
		log.Printf("%s is moved to %s", oldLoc, newLoc)
	}
	return nil
}

func checkDestinationDir(c *config) error {
	_, err := ioutil.ReadDir(c.Dest)
	return err
}
