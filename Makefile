SHELL=/bin/bash

build_wasm:
	GOOS=js GOARCH=wasm go build -o ./public/vita.wasm .

copy_wasm_exec:
	cp "`go env GOROOT`/misc/wasm/wasm_exec.js" ./public/

build: build_wasm copy_wasm_exec

run: build
	serve ./public

build_wasm_tinygo:
	tinygo build -o ./public/vita.wasm -target wasm .

copy_wasm_exec_tinygo:
	cp "`tinygo env TINYGOROOT`/targets/wasm_exec.js" ./public/

build_tinygo: build_wasm_tinygo copy_wasm_exec_tinygo
