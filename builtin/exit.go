package builtin

import (
	"os"
	"strings"
)

type Exit struct {
}

func (e *Exit) Handle(args []string) {
	if len(args) > 0 && strings.TrimSpace(args[0]) == "0" {
		os.Exit(0)
	}
}

func (e *Exit) GetCmdName() string {
	return "exit"
}
