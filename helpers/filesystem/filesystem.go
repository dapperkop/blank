package filesystem

import (
	"os"

	"github.com/dapperkop/blank/logger"
)

// CreateDir func ...
func CreateDir(path string) {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		logger.Logger.Fatalln(err)
	}
}

// Getwd func ...
func Getwd() string {
	var (
		dir string
		err error
	)

	dir, err = os.Getwd()

	if err != nil {
		logger.Logger.Fatalln(err)
	}

	return dir
}

// IsExist func ...
func IsExist(path string) bool {
	var err error

	_, err = os.Stat(path)

	return !os.IsNotExist(err)
}

// IsNotExist func ...
func IsNotExist(path string) bool {
	var err error

	_, err = os.Stat(path)

	return os.IsNotExist(err)
}

// Rename func ...
func Rename(oldpath, newpath string) {
	if err := os.Rename(oldpath, newpath); err != nil {
		logger.Logger.Fatalln(err)
	}
}
