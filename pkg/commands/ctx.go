package commands

import (
	"context"
)

type key int

const (
	ConsoleCtx key = iota
)

type Ctx struct {
	cwd string
}

func NewContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ConsoleCtx, &Ctx{cwd: ""})
}

func MustFromContext(ctx context.Context) *Ctx {
	u, ok := ctx.Value(ConsoleCtx).(*Ctx)
	if !ok {
		panic("invalid context")
	}

	return u
}

func FromContext(ctx context.Context) (*Ctx, bool) {
	u, ok := ctx.Value(ConsoleCtx).(*Ctx)
	return u, ok
}
