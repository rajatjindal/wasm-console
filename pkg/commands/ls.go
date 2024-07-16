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
	if (len(args) == 1 && GetPwd(ctx) == rootDir(ctx)) || (len(args) > 1 && args[1] == rootDir(ctx)) {
		lsRoot(ctx)
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

func lsRoot(ctx context.Context) {
	if MustFromContext(ctx).runtime == RuntimeSpin {
		entries, err := os.ReadDir(rootDir(ctx))
		if err != nil {
			fmt.Printf("failed to read root dir. error: %v\n", err)
			return
		}

		for _, e := range entries {
			fmt.Println(e.Name())
		}

		return
	}

	for _, d := range preopens.GetDirectories().Slice() {
		fmt.Println(d.F1)
	}
}
