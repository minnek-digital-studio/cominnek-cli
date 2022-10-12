package files

import "os"

func CheckExist(filename string) bool {
	_, err := os.Stat(filename)

	return !os.IsNotExist(err)
}