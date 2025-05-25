package builtin

import (
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/executable"
	"strings"
)

type Type struct {
}

func (e *Type) Handle(args []string) {
	command := strings.TrimSpace(args[0])

	value, exists := GetBuiltInMap()[command]

	if exists {
		fmt.Printf("%s is a shell builtin\n", value.GetCmdName())
		return
	}

	if file, exists := executable.IsInPath(command); exists {
		fmt.Printf("%s is %s\n", command, file)
		return
	}

	fmt.Printf("%s: not found\n", strings.TrimSpace(args[0]))
}

func (e *Type) GetCmdName() string {
	return "type"
}
