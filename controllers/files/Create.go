package files

import (
	"io/ioutil"
	"log"
)

func Create(content []byte, fileName string) {
	err := ioutil.WriteFile(fileName, content, 0777)

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
}