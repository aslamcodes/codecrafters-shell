package builtin

import (
	"github.com/codecrafters-io/shell-starter-go/utils"
)

type Pwd struct {
}

func (e *Pwd) Handle(args []string) {
	utils.GetCurrPath()
}

func (e *Pwd) GetCmdName() string {
	return "pwd"
}
