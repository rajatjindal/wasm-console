package commands

import (
	"fmt"
	"os"
)

func Rm(args []string) error {
	if len(args) <= 0 {
		return fmt.Errorf("atleast one arg is required for rm")
	}

	//TODO: check if dir or file
	return os.Remove(args[0])
}
