Experimental wasm-console, because why not?

### Show me aleady

```
$ /usr/local/bin/wasmtime run --dir testme --dir pkg -Shttp --env ABC=XYZ  main.wasm 
starting wasm-console. Enter 'exit' to quit the shell.
wasm-console > printenv
ABC=XYZ
wasm-console > pwd
/
wasm-console > ls
testme
pkg
wasm-console > cd testme
wasm-console testme> ls
some-dir
some-file
wasm-console testme> cd some-dir
wasm-console testme/some-dir> ls
some-other-file
wasm-console testme/some-dir> cat some-other-file
contents of some-other-file
wasm-console testme/some-dir> cd ../ 
wasm-console testme> ls
some-dir
some-file
wasm-console testme> cat some-file
some - content
wasm-console testme> curl https://google.com
HTTP/0.0 301 Moved Permanently

<HTML><HEAD><meta http-equiv="content-type" content="text/html;charset=utf-8">
<TITLE>301 Moved</TITLE></HEAD><BODY>
<H1>301 Moved</H1>
The document has moved
<A HREF="https://www.google.com/">here</A>.
</BODY></HTML>

wasm-console testme> history
1* printenv
2* pwd
3* ls
4* cd testme
5* ls
6* cd some-dir
7* ls
8* cat some-other-file
9* cd ../
10* ls
11* cat some-file
12* curl https://google.com
13* history
wasm-console > ^D
goodbye !
```
## As a gif
![using wasm-console](demo/usage.gif)

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
