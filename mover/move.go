package mover

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

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
