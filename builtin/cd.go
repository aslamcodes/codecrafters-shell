package builtin

import (
	"fmt"
	"os"
)

type Cd struct {
}

func (e *Cd) Handle(args []string) {
	dir := args[0]

	entries, err := os.ReadDir("./")

	if err != nil {
		fmt.Println("Cannot access directories")
		return
	}

	for _, e := range entries {
		if dir == e.Name() {
			os.Chdir(args[0])
			return
		}
	}

	fmt.Printf("cd: %s: No such file or directory\n", args[0])
}

func (e *Cd) GetCmdName() string {
	return "cd"
}
