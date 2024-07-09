Experimental wasm-console, because why not?

## Building
```
$ ls
README.md build.sh  go.mod    go.sum    internal  main.go   pkg       testme    wit

## build wasm component
$ ./build.sh

## ls, notice main.wasm file is created
$ ls
README.md build.sh  go.mod    go.sum    internal  main.go   main.wasm pkg       testme    wit
```


## Running
```
$ wasmtime run --dir testme main.wasm
starting wasm-console. Enter 'exit' to quit the shell.
wasm-console> ls testme
some-dir
some-file
wasm-console> ls testme/some-dir
some-other-file
wasm-console> cat testme/some-dir/some-other-file
contents of some-other-file
wasm-console> curl https://random-data-api.fermyon.app/animals/json
HTTP/0.0 200 OK

{"timestamp":1720316222127,"fact":"Reindeer grow new antlers every year"}
wasm-console> history
1* ls testme
2* ls testme/some-dir
3* cat testme/some-dir/some-other-file
4* curl https://random-data-api.fermyon.app/animals/json
5* history
wasm-console> exit
```
