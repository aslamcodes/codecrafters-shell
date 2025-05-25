package utils

import (
	"os"
)

func GetCurrPath() (string, error) {
	mydir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return mydir, nil
}

func IsValidPath(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
