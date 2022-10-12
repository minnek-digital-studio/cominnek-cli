package folders

import "os"

func Create(_path string) {
	os.Mkdir(_path, os.ModePerm)
}