package builtin

import (
	"fmt"
	"strings"
)

type Echo struct {
}

func (e *Echo) Handle(args []string) {
	fmt.Println(strings.TrimSpace(strings.Join(args, " ")))
}

func (e *Echo) GetCmdName() string {
	return "echo"
}
