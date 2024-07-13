package commands

import (
	"context"
	"fmt"
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

	MustFromContext(ctx).cwd = dir
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
