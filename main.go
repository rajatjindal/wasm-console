package main

import (
	"fmt"
	"strings"

	"github.com/rajatjindal/wasmshell/internal/wasi/cli/stdin"
	"github.com/rajatjindal/wasmshell/pkg/commands"
)

func main() {
	history := []string{}
	fmt.Println("starting wasmshell. Enter 'exit' to quit the shell.")
	for {
		fmt.Print("wasmshell> ")
		result := stdin.GetStdin().BlockingRead(1024)
		if result.IsErr() {
			if result.Err().Closed() {
				fmt.Println("\ngoodbye !")
				return
			}
			panic("wasmshell error")
		}

		input := strings.TrimSpace(string(result.OK().Slice()))
		cmdAndArgs := strings.Split(input, " ")
		if len(cmdAndArgs) == 0 {
			fmt.Println("no command provided")
			continue
		}

		history = append(history, string(result.OK().Slice()))
		cmd := strings.TrimSpace((cmdAndArgs[0]))
		switch cmd {
		case "ls":
			err := commands.Ls(cmdAndArgs[1:])
			if err != nil {
				fmt.Println(err)
			}
		case "rm":
			err := commands.Rm(cmdAndArgs[1:])
			if err != nil {
				fmt.Println(err)
			}
		case "cat":
			err := commands.Cat(cmdAndArgs[1:])
			if err != nil {
				fmt.Println(err)
			}
		case "history":
			for i, h := range history {
				fmt.Print(i+1, "* ", h)
			}
		case "curl":
			err := commands.Curl(cmdAndArgs[1:])
			if err != nil {
				fmt.Println(err)
			}
		case "exit":
			return
		case "":
			continue
		default:
			fmt.Printf("command %q not supported\n", cmd)
		}
	}
}
