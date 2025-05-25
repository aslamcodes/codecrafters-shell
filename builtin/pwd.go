package builtin

import (
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/utils"
)

type Pwd struct {
}

func (e *Pwd) Handle(args []string) {
	dir, _ := utils.GetCurrPath()
	fmt.Println(dir)
}

func (e *Pwd) GetCmdName() string {
	return "pwd"
}
