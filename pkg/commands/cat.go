package commands

import (
	"context"
	"fmt"
	"os"
	"path"
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

func isAbsolute(inp string) bool {
	return strings.HasPrefix(inp, "/")
}

func getAbsolute(ctx context.Context, inp string) string {
	if isAbsolute(inp) {
		return inp
	}

	return path.Join(GetPwd(ctx), inp)
}
