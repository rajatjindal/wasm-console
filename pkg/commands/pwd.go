package commands

import (
	"context"
	"fmt"
)

func Pwd(ctx context.Context) {
	//always prefix "/" when printing pwd
	if needsRootPrefix(ctx) {
		fmt.Println("/" + GetPwd(ctx))
		return
	}

	fmt.Println(GetPwd(ctx))
}

func GetPwd(ctx context.Context) string {
	return MustFromContext(ctx).cwd
}
