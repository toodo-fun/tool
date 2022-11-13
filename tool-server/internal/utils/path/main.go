package path

import "os"

func IsExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {

		return false
	}
	return true
}

func MkdirAll(path string) (err error) {
	err = os.MkdirAll(path, 655)
	return err
}
