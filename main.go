package main

import (
	"math"
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
	universe    *Universe
	ctx         js.Value
	animationID int = -1
)

func main() {
	done := make(chan bool)

	universe = NewUniverse()
	window := js.Global()
	document := window.Get("document")

	canvas := setupCanvas()
	ctx = canvas.Call("getContext", "2d")

	// Rendering loop
	var draw js.Func
	draw = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		universe.tick()
		drawCanvas()
		animationID = window.Call("requestAnimationFrame", draw).Int()
		return nil
	})
	defer draw.Release()

	// Add event listeners

	canvasClickListener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		boundingRect := canvas.Call("getBoundingClientRect")
		widthScale := canvas.Get("width").Int() / boundingRect.Get("width").Int()
		heightScale := canvas.Get("height").Int() / boundingRect.Get("height").Int()
		canvasX := (args[0].Get("clientX").Int() - boundingRect.Get("left").Int()) * widthScale
		canvasY := (args[0].Get("clientY").Int() - boundingRect.Get("top").Int()) * heightScale
		row := math.Floor(float64(canvasY) / (cellSize + borderSize))
		col := math.Floor(float64(canvasX) / (cellSize + borderSize))

		universe.toggleCellAt(uint32(row), uint32(col))
		drawCanvas()
		return nil
	})
	defer canvasClickListener.Release()
	canvas.Call("addEventListener", "click", canvasClickListener)

	playPauseButton := document.Call("getElementById", "play-pause")
	playPauseListener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if animationID != -1 {
			playPauseButton.Set("textContent", "Play")
			window.Call("cancelAnimationFrame", animationID)
			animationID = -1
		} else {
			playPauseButton.Set("textContent", "Pause")
			window.Call("requestAnimationFrame", draw)
		}
		return nil
	})
	defer playPauseListener.Release()
	playPauseButton.Call("addEventListener", "click", playPauseListener)

	resetButton := document.Call("getElementById", "reset")
	resetListener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		universe.reset()
		drawCanvas()
		return nil
	})
	defer resetListener.Release()
	resetButton.Call("addEventListener", "click", resetListener)

	randomizeButton := document.Call("getElementById", "randomize")
	randomizeListener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		universe = NewUniverse()
		drawCanvas()
		return nil
	})
	defer randomizeListener.Release()
	randomizeButton.Call("addEventListener", "click", randomizeListener)

	// Start rendering
	window.Call("requestAnimationFrame", draw)

	// Block and wait for event listeners
	<-done
}

func setupCanvas() js.Value {
	document := js.Global().Get("document")
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