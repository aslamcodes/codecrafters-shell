package shell

import "github.com/codecrafters-io/shell-starter-go/utils"

func (s *Shell) Init() {
	s.WD, _ = utils.GetCurrPath()
}
