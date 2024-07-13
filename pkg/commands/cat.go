package commands

import (
	"context"
	"fmt"
	"os"
	"strings"
)

func Cat(ctx context.Context, input string) error {
	args := strings.Fields(strings.TrimSpace(input))
	if len(args) <= 1 {
		return fmt.Errorf("atleast one arg is required for `cat`")
	}

	file := getAbsolute(ctx, args[1])
	raw, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	fmt.Println(string(raw))
	return nil
}
