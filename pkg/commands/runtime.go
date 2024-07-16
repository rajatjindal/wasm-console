package commands

import (
	"context"
	"path"
	"strings"
)

// TODO(rajatjindal): figure out why this difference and fix accordingly
// wasmtime does not like absolute paths, and
// spin does not like relative paths.
func rootDir(ctx context.Context) string {
	if MustFromContext(ctx).runtime == RuntimeSpin {
		return "/"
	}

	return ""
}

func needsRootPrefix(ctx context.Context) bool {
	return MustFromContext(ctx).runtime == RuntimeSpin
}

func isAbsolute(inp string) bool {
	return strings.HasPrefix(inp, "/")
}

func getAbsolute(ctx context.Context, inp string) string {
	abs := getAbsoluteInner(ctx, inp)
	switch {
	case needsRootPrefix(ctx) && !strings.HasPrefix(abs, "/"):
		return "/" + abs
	case !needsRootPrefix(ctx) && strings.HasPrefix(abs, "/"):
		return strings.TrimPrefix(abs, "/")
	}

	return abs
}

func getAbsoluteInner(ctx context.Context, inp string) string {
	if isAbsolute(inp) {
		return inp
	}

	return path.Join(GetPwd(ctx), inp)
}
