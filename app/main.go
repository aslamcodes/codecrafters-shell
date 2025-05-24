package main

import (
	"bufio"
	"fmt"
	"os"
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
			value, exists := builtin[strings.TrimSpace(args[0])]

			if exists {
				fmt.Printf("%s is a shell builtin\n", value)
				break
			}

			path := os.Getenv("PATH")

			if path == "" {
				break
			}

			found := false
			exec := strings.TrimSpace(args[0])
			for location := range strings.SplitSeq(path, string(os.PathListSeparator)) {
				file := location + "/" + exec
				if _, err := os.Stat(file); err == nil {
					fmt.Printf("%s is %s\n", exec, file)
					found = true
					break
				}
			}

			if found {
				break
			}

			fmt.Printf("%s: not found\n", strings.TrimSpace(args[0]))
		default:
			fmt.Println(command[:len(command)-1] + ": command not found")
		}
	}
}
