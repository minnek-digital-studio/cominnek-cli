package folders

import "os"

func Create(_path string) {
	err := os.MkdirAll(_path, 0755)

	if err != nil {
		panic(err)
	}
}
