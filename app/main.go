package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
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

		case "exit":
			if strings.TrimSpace(args[0]) == "0" {
				os.Exit(0)
			}
		case "echo":
			fmt.Println(strings.TrimSpace(strings.Join(args, " ")))
		default:
			fmt.Println(command[:len(command)-1] + ": command not found")
		}
	}
}
