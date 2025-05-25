package builtin

import (
	"os"
)

type Cd struct {
}

func (e *Cd) Handle(args []string) {
	os.Chdir(args[0])
}

func (e *Cd) GetCmdName() string {
	return "cd"
}
