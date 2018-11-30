package mover

import (
	"log"
	"os"
	"time"
)

func Controller() {

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

		time.Sleep(time.Second * 10)
	}

}
