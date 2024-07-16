package commands

import (
	"context"
)

type key int

const (
	ConsoleCtx key = iota
)

type Runtime string

const (
	RuntimeWasmtime Runtime = "wasmtime"
	RuntimeSpin     Runtime = "spin"
)

type Ctx struct {
	cwd     string
	runtime Runtime
}

func NewContext(ctx context.Context) context.Context {
	runtime := guessRuntime()
	rootdir := ""
	if runtime == RuntimeSpin {
		rootdir = "/"
	}

	return context.WithValue(ctx, ConsoleCtx, &Ctx{cwd: rootdir, runtime: runtime})
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

func guessRuntime() Runtime {
	// if preopens.GetDirectories().Len() == 1 && preopens.GetDirectories().Slice()[0].F1 == "/" {
	// 	return RuntimeSpin
	// }

	// return RuntimeWasmtime
	return RuntimeSpin
}
