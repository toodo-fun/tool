package path

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

var (
	runModel = ""
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
	if runModel == "DEBUG" {
		return "C:\\Users\\marui\\Documents\\project\\tool\\tool-server"
	} else {
		return path
	}
}

func init() {
	if !IsExist(".env") {
		os.Create(".env")
	}

	err := godotenv.Load()
	if err != nil {
		logrus.Fatalln(err)
	}

	runModel = os.Getenv("RUN_MODEL")
}
