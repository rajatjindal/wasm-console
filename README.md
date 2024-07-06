Experimental wasmshell, because why not?

```
$ ls
README.md build.sh  go.mod    go.sum    internal  main.go   pkg       testme    wit

## build wasm component
$ ./build.sh

## ls, notice main.wasm file is created
$ ls
README.md build.sh  go.mod    go.sum    internal  main.go   main.wasm pkg       testme    wit

## Run using wasmtime
$ wasmtime run --dir testme main.wasm
starting wasmshell. Enter 'exit' to quit the shell.
wasmshell> ls testme
some-dir
some-file
wasmshell> ls testme/some-dir
some-other-file
wasmshell> cat testme/some-dir/some-other-file
contents of some-other-file
wasmshell> history
1* ls testme
2* ls testme/some-dir
3* cat testme/some-dir/some-other-file
4* history
wasmshell> exit
```