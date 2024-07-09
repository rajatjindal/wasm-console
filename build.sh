## build with debug logs from tinygo
# tinygo build -x -target=wasip2 --wit-package $(go list -mod=readonly -m -f '{{.Dir}}' github.com/rajatjindal/wasm-console)/wit --wit-world cli -o main.wasm -x -work main.go

## or without them
tinygo build -target=wasip2 --wit-package $(go list -mod=readonly -m -f '{{.Dir}}' github.com/rajatjindal/wasm-console)/wit --wit-world cli -o main.wasm main.go
