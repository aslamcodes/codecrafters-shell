package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	builtin := map[string]string{
		"echo": "echo",
		"exit": "exit",
		"type": "type",
	}

	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		tokens := strings.Split(command, " ")
		program := tokens[0]
		args := tokens[1:]

		switch strings.TrimSpace(program) {
		case builtin["exit"]:
			if strings.TrimSpace(args[0]) == "0" {
				os.Exit(0)
			}
		case builtin["echo"]:
			fmt.Println(strings.TrimSpace(strings.Join(args, " ")))
		case builtin["type"]:
			command := strings.TrimSpace(args[0])

			value, exists := builtin[command]

			if exists {
				fmt.Printf("%s is a shell builtin\n", value)
				break
			}

			if file, exists := isInPath(command); exists {
				fmt.Printf("%s is %s\n", command, file)
				break
			}

			fmt.Printf("%s: not found\n", strings.TrimSpace(args[0]))

		default:
			if file, exists := isInPath(strings.TrimSpace(program)); exists {
				cmd := exec.Command(file, strings.Join(args, " "))
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()
				break
			}
			fmt.Println(command[:len(command)-1] + ": command not found")
		}
	}
}

func isInPath(exec string) (string, bool) {
	path := os.Getenv("PATH")
	for location := range strings.SplitSeq(path, string(os.PathListSeparator)) {
		file := location + "/" + exec
		if _, err := os.Stat(file); err == nil {
			return file, true
		}
	}
	return "", false
}
