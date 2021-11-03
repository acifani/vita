SHELL=/bin/bash

build_wasm:
	GOOS=js GOARCH=wasm go build -o ./public/vita.wasm .

copy_wasm_exec:
	cp "`go env GOROOT`/misc/wasm/wasm_exec.js" ./public/

build: build_wasm copy_wasm_exec

run: build
	serve ./public
