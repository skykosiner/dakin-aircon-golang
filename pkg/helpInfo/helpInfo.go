package helpinfo

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func doesFileExist() bool {
	path := fmt.Sprintf("%s/.local/airconhelp.txt", os.Getenv("HOME"))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func MoveHelpFile() {
	// If the help file does not exist create the help file in the correct location
	if !doesFileExist() {
		src := "./helptext.txt"
		dest := fmt.Sprintf("%s/.local/airconhelp.txt", os.Getenv("HOME"))

		bytesRead, err := ioutil.ReadFile(src)

		if err != nil {
			log.Fatal("Error setting up help file ", err)
		}

		file, err := os.Create(dest)

		if err != nil {
			log.Fatal("Error setting up help file ", err, file)
		}

		err = ioutil.WriteFile(dest, bytesRead, 0644)

		if err != nil {
			log.Fatal("Error setting up help file ", err)
		}
	}
}
