package commands

import (
	"fmt"
	"os"
)

func Ls(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("cannot list under / dir. pls provide the dir name mounted using wasmtime")
	}

	dir := args[0]
	//TODO: check if dir or file
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fmt.Println(entry.Name())
	}

	return nil
}
