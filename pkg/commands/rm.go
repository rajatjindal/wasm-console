package commands

import (
	"fmt"
	"os"
	"strings"
)

func Rm(input string) error {
	args := strings.Fields(strings.TrimSpace(input))
	if len(args) == 1 {
		return fmt.Errorf("atleast one arg is required for `rm`")
	}

	//TODO: check if dir or file
	return os.Remove(args[1])
}
