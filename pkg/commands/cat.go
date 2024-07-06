package commands

import (
	"fmt"
	"os"
)

func Cat(args []string) error {
	if len(args) <= 0 {
		return fmt.Errorf("atleast one arg is required for cat")
	}

	raw, err := os.ReadFile(args[0])
	if err != nil {
		return err
	}

	fmt.Println(string(raw))
	return nil
}
