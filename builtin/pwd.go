package builtin

import (
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/utils"
)

type Pwd struct {
}

func (e *Pwd) Handle(args []string) {
	fmt.Println(utils.GetCurrPath())
}

func (e *Pwd) GetCmdName() string {
	return "pwd"
}
