package commands

import (
	"context"
	"fmt"
	"os"
)

func PrintEnv(ctx context.Context) error {
	for _, v := range os.Environ() {
		fmt.Println(v)
	}

	return nil
}
