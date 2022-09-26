package files

import (
	"io/ioutil"
	"log"
	"os"
)

func Read(fileName string) []byte {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return []byte("")
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	return data
}