package path

import (
	"os"
	"path/filepath"
)

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

func GetPlatformRoot() string {
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return path
}
