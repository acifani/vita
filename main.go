package main

import (
	"syscall/js"
)

const (
	dead  = iota
	alive = iota
)

const (
	cellSize   = 10
	borderSize = 1
)

var (
	universe *Universe
	ctx      js.Value
)

func main() {
	done := make(chan bool)

	universe = NewUniverse()
	window := js.Global()
	canvas := setupCanvas()
	ctx = canvas.Call("getContext", "2d")

	var draw js.Func
	draw = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		universe.tick()
		drawCanvas()
		window.Call("requestAnimationFrame", draw)
		return nil
	})
	defer draw.Release()

	window.Call("requestAnimationFrame", draw)

	<-done
}

func setupCanvas() js.Value {
	window := js.Global()
	document := window.Get("document")
	canvas := document.Call("getElementById", "canvas")
	canvas.Set("height", (cellSize+borderSize)*universe.height+borderSize)
	canvas.Set("width", (cellSize+borderSize)*universe.width+borderSize)

	return canvas
}

func drawCanvas() {
	drawGrid()
	drawCells()
}

func drawGrid() {
	height := int(universe.height)
	width := int(universe.width)

	ctx.Call("beginPath")
	ctx.Set("strokeStyle", "#aaa")

	for i := 0; i <= width; i++ {
		ctx.Call("moveTo", i*(cellSize+borderSize)+borderSize, 0)
		ctx.Call("lineTo", i*(cellSize+borderSize)+borderSize, (cellSize+borderSize)*height+borderSize)
	}

	for j := 0; j <= height; j++ {
		ctx.Call("moveTo", 0, j*(cellSize+borderSize)+borderSize)
		ctx.Call("lineTo", (cellSize+borderSize)*width+borderSize, j*(cellSize+borderSize)+borderSize)
	}

	ctx.Call("stroke")
}

func drawCells() {
	height := int(universe.height)
	width := int(universe.width)

	ctx.Call("beginPath")

	// Live cells
	ctx.Set("fillStyle", "#222")
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			idx := universe.getIndex(uint32(row), uint32(col))
			if universe.cells[idx] == alive {
				ctx.Call("fillRect",
					col*(cellSize+borderSize)+borderSize,
					row*(cellSize+borderSize)+borderSize,
					cellSize,
					cellSize,
				)
			}
		}
	}

	// Dead cells
	ctx.Set("fillStyle", "#fafafa")
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			idx := universe.getIndex(uint32(row), uint32(col))
			if universe.cells[idx] == dead {
				ctx.Call("fillRect",
					col*(cellSize+borderSize)+borderSize,
					row*(cellSize+borderSize)+borderSize,
					cellSize,
					cellSize,
				)
			}
		}
	}

	ctx.Call("stroke")
}
