package commands

import (
	"context"
	"fmt"
	"os"
	"strings"
)

func Rm(ctx context.Context, input string) error {
	args := strings.Fields(strings.TrimSpace(input))
	if len(args) == 1 {
		return fmt.Errorf("atleast one arg is required for `rm`")
	}

	file := getAbsolute(ctx, args[1])
	return os.Remove(file)
}
