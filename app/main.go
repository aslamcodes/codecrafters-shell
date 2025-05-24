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

		if strings.TrimSpace(program) == "exit" && strings.TrimSpace(args[0]) == "0" {
			os.Exit(0)
		}

		fmt.Println(command[:len(command)-1] + ": command not found")
	}
}
