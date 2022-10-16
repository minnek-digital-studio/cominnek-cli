package files

import "os"

func Delete(file string) {
	os.Remove(file)
}