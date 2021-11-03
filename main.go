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
	universeString := universe.toString()

	global := js.Global()
	document := global.Get("document")

	canvas := document.Call("getElementById", "canvas")
	canvas.Set("innerText", universeString)
	<-done
}
