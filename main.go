package main

import (
	"syscall/js"
)

const (
	dead  = iota
	alive = iota
)

func main() {
	done := make(chan bool)

	universe := NewUniverse()

	window := js.Global()
	document := window.Get("document")
	canvas := document.Call("getElementById", "canvas")

	var draw js.Func
	draw = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		canvas.Set("innerText", universe.toString())
		universe.tick()
		window.Call("requestAnimationFrame", draw)
		return nil
	})

	window.Call("requestAnimationFrame", draw)

	<-done
}
