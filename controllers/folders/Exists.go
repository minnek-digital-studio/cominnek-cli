package folders

import "os"

func CheckExists(_path string) bool {
	if _, err := os.Stat(_path); os.IsNotExist(err) {
		return false;
	}

	return true;
}