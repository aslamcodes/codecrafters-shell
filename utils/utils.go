package utils

import (
	"fmt"
	"os"
)

func GetCurrPath() {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mydir)
}
