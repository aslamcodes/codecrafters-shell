package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetCurrPath() {
	ex, err := os.Executable()

	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
}
