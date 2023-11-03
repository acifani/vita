SHELL=/bin/bash

test:
	go test ./lib/game/ -v

tinygo_test:
	tinygo test -v ./lib/game/

build_wasm: test
	GOOS=js GOARCH=wasm go build -o ./public/vita.wasm .

copy_wasm_exec:
	cp "`go env GOROOT`/misc/wasm/wasm_exec.js" ./public/

build: build_wasm copy_wasm_exec

run: build
	go run ./serve

run_tinygo: build_tinygo
	go run ./serve

build_wasm_tinygo: tinygo_test
	tinygo build -o ./public/vita.wasm -target wasm -no-debug -panic=trap -opt=s .

copy_wasm_exec_tinygo:
	cp "`tinygo env TINYGOROOT`/targets/wasm_exec.js" ./public/

build_tinygo: build_wasm_tinygo copy_wasm_exec_tinygo
