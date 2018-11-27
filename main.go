package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v\n", err)
	}
	defer f.Close()

	log.SetOutput(f)

	log.Printf("File mover started..\n")
	c, err := readConfig()
	if err != nil {
		log.Fatal("error while reading configuration file")
	}

	for {
		if err := moveFile(c); err != nil {
			log.Fatal(err.Error() + "\n")
		}

		time.Sleep(time.Second * 100)
	}

}

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

func moveFile(c *config) error {

	log.Printf("Started to moving files\n")
	files, err := ioutil.ReadDir(c.Source)
	if err != nil {
		return errors.New("Source destination location doesnt exists " + err.Error())
	}
	if err := checkDestinationDir(c); err != nil {
		return errors.New("Destination location doest exists " + err.Error())
	}
	if len(files) > 0 {
		oldLoc := c.Source + "/" + files[0].Name()
		newLoc := c.Dest + "/" + c.File
		err := os.Rename(oldLoc, newLoc)
		if err != nil {
			return errors.New("Failed to move files " + err.Error())
		}
		log.Printf("%s is moved to %s\n", oldLoc, newLoc)
	}
	return nil
}

func checkDestinationDir(c *config) error {
	_, err := ioutil.ReadDir(c.Dest)
	return err
}

func fileRefractor(filename string) error {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	data := string(f)
	newData := strings.Replace(data, "\\", "/", 10)
	fmt.Println(newData)
	return nil
}
