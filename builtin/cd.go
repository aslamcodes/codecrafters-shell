package builtin

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/utils"
)

type Cd struct {
}

func (e *Cd) Handle(args []string) {
	dir := args[0]

	if dir == "~" {
		dir = os.Getenv("HOME")
	}

	if !utils.IsValidPath(dir) {
		fmt.Printf("cd: %s: No such file or directory\n", args[0])
	}

	os.Chdir(dir)
}

func (e *Cd) GetCmdName() string {
	return "cd"
}
