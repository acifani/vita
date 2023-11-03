SHELL=/bin/bash

test:
	tinygo test -v ./lib/game/

build_wasm: test
	tinygo build -o ./public/vita.wasm -target wasm -no-debug -panic=trap -opt=s .

copy_wasm_exec:
	cp "`tinygo env TINYGOROOT`/targets/wasm_exec.js" ./public/

build: build_wasm copy_wasm_exec

run: build
	go run ./serve

build_wasi:
	spin build

run_wasi: build_wasi
	spin up
