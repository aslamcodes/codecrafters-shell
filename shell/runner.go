package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/builtin"
	"github.com/codecrafters-io/shell-starter-go/executable"
)

func (s *Shell) RunShell() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		prompt, err := bufio.NewReader(os.Stdin).ReadString('\n')
		prompt = strings.TrimSpace(prompt)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		program, args := parsePrompt(prompt)
		builtin, exists := builtin.GetBuiltInMap()[program]

		if exists {
			builtin.Handle(args)
			continue
		}

		_, exists = executable.IsInPath(program)

		if exists {
			executable.RunExecutable(program, args)
			continue
		}

		fmt.Println(program + ": command not found")
	}
}

func parsePrompt(prompt string) (string, []string) {
	tokens := strings.Split(prompt, " ")
	program := tokens[0]
	args := tokens[1:]
	return program, args
}
