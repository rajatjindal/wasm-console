package commands

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/rajatjindal/wasm-console/internal/wasi/filesystem/preopens"
)

func Ls(ctx context.Context, input string) error {
	args := strings.Fields(strings.TrimSpace(input))
	if (len(args) == 1 && GetPwd(ctx) == "") || (len(args) > 1 && strings.TrimPrefix(args[1], "/") == "") {
		lsRoot()
		return nil
	}

	dir := GetPwd(ctx)
	if len(args) > 1 {
		dir = getAbsolute(ctx, args[1])
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fmt.Println(entry.Name())
	}

	return nil
}

func lsRoot() {
	for _, d := range preopens.GetDirectories().Slice() {
		fmt.Println(d.F1)
	}
}
