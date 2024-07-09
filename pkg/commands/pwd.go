package commands

import (
	"context"
	"fmt"
)

func Pwd(ctx context.Context) {
	fmt.Println("/" + GetPwd(ctx))
}

func GetPwd(ctx context.Context) string {
	return MustFromContext(ctx).cwd
}
