package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/rajatjindal/wasm-console/internal/wasi/cli/stdin"
	"github.com/rajatjindal/wasm-console/pkg/commands"
)

func main() {
	ctx := commands.NewContext(context.Background())
	history := []string{}
	fmt.Println("starting wasm-console. Enter 'exit' to quit the shell.")
	for {
		fmt.Printf("wasm-console %s> ", commands.GetPwd(ctx))
		result := stdin.GetStdin().BlockingRead(1024)
		if result.IsErr() {
			if result.Err().Closed() {
				fmt.Println("\ngoodbye !")
				return
			}
			panic("wasm-console error")
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
		case "cd":
			commands.Cd(ctx, input)
		case "pwd":
			commands.Pwd(ctx)
		case "printenv":
			commands.PrintEnv(ctx)
		case "ls":
			err := commands.Ls(ctx, input)
			if err != nil {
				fmt.Println(err)
			}
		case "rm":
			err := commands.Rm(ctx, input)
			if err != nil {
				fmt.Println(err)
			}
		case "cat":
			err := commands.Cat(ctx, input)
			if err != nil {
				fmt.Println(err)
			}
		case "history":
			for i, h := range history {
				fmt.Print(i+1, "* ", h)
			}
		case "curl":
			err := commands.Curl(input)
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
