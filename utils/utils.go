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
