package commands

import (
	"context"
	"fmt"
	"os"
	"path"
	"strings"
)

func Cd(ctx context.Context, input string) error {
	args := strings.Fields(strings.TrimSpace(input))
	if len(args) < 2 {
		MustFromContext(ctx).cwd = ""
		return nil
	}

	dir, err := calculateDir(ctx, args[1])
	if err != nil {
		return err
	}

	woutPrefix := strings.TrimPrefix(dir, "/")
	_, err = os.Stat(woutPrefix)
	if err != nil {
		return fmt.Errorf("cd: no such file or dir: %s", args[1])
	}

	MustFromContext(ctx).cwd = woutPrefix
	return nil
}

func calculateDir(ctx context.Context, dir string) (string, error) {
	if isAbsolute(dir) {
		return dir, nil
	}

	if strings.HasPrefix(dir, "..") {
		cwd := GetPwd(ctx)
		parts := strings.Split(dir, "/")
		for _, p := range parts {
			if p == ".." {
				cwdParts := strings.Split(cwd, "/")
				if cwd == "" || len(cwdParts) == 0 {
					return "", fmt.Errorf("cd: no such file or directory: %s", dir)
				}

				cwd = strings.Join(cwdParts[:len(cwdParts)-1], "/")
			} else {
				cwd = path.Join(cwd, p)
			}
		}

		return cwd, nil
	}

	//relative path
	return path.Join(GetPwd(ctx), dir), nil
}
