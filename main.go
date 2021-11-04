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

	addEventListener("canvas", "click", func(this js.Value, args []js.Value) interface{} {
		boundingRect := canvas.Call("getBoundingClientRect")
		widthScale := canvas.Get("width").Int() / boundingRect.Get("width").Int()
		heightScale := canvas.Get("height").Int() / boundingRect.Get("height").Int()
		canvasX := (args[0].Get("clientX").Int() - boundingRect.Get("left").Int()) * widthScale
		canvasY := (args[0].Get("clientY").Int() - boundingRect.Get("top").Int()) * heightScale
		row := uint32(math.Floor(float64(canvasY) / (cellSize + borderSize)))
		col := uint32(math.Floor(float64(canvasX) / (cellSize + borderSize)))

		if args[0].Get("ctrlKey").Bool() {
			// Insert glider
			universe.toggleCellAt(row-1, col)
			universe.toggleCellAt(row, col+1)
			universe.toggleCellAt(row+1, col-1)
			universe.toggleCellAt(row+1, col)
			universe.toggleCellAt(row+1, col+1)
		} else if args[0].Get("shiftKey").Bool() {
			// Insert pulsar
			universe.toggleCellAt(row-6, col-4)
			universe.toggleCellAt(row-6, col-3)
			universe.toggleCellAt(row-6, col-2)
			universe.toggleCellAt(row-6, col+2)
			universe.toggleCellAt(row-6, col+3)
			universe.toggleCellAt(row-6, col+4)

			universe.toggleCellAt(row-4, col-6)
			universe.toggleCellAt(row-4, col-1)
			universe.toggleCellAt(row-4, col+1)
			universe.toggleCellAt(row-4, col+6)

			universe.toggleCellAt(row-3, col-6)
			universe.toggleCellAt(row-3, col-1)
			universe.toggleCellAt(row-3, col+1)
			universe.toggleCellAt(row-3, col+6)

			universe.toggleCellAt(row-2, col-6)
			universe.toggleCellAt(row-2, col-1)
			universe.toggleCellAt(row-2, col+1)
			universe.toggleCellAt(row-2, col+6)

			universe.toggleCellAt(row-1, col-4)
			universe.toggleCellAt(row-1, col-3)
			universe.toggleCellAt(row-1, col-2)
			universe.toggleCellAt(row-1, col+2)
			universe.toggleCellAt(row-1, col+3)
			universe.toggleCellAt(row-1, col+4)

			universe.toggleCellAt(row+1, col-4)
			universe.toggleCellAt(row+1, col-3)
			universe.toggleCellAt(row+1, col-2)
			universe.toggleCellAt(row+1, col+2)
			universe.toggleCellAt(row+1, col+3)
			universe.toggleCellAt(row+1, col+4)

			universe.toggleCellAt(row+2, col-6)
			universe.toggleCellAt(row+2, col-1)
			universe.toggleCellAt(row+2, col+1)
			universe.toggleCellAt(row+2, col+6)

			universe.toggleCellAt(row+3, col-6)
			universe.toggleCellAt(row+3, col-1)
			universe.toggleCellAt(row+3, col+1)
			universe.toggleCellAt(row+3, col+6)

			universe.toggleCellAt(row+4, col-6)
			universe.toggleCellAt(row+4, col-1)
			universe.toggleCellAt(row+4, col+1)
			universe.toggleCellAt(row+4, col+6)

			universe.toggleCellAt(row+6, col-4)
			universe.toggleCellAt(row+6, col-3)
			universe.toggleCellAt(row+6, col-2)
			universe.toggleCellAt(row+6, col+2)
			universe.toggleCellAt(row+6, col+3)
			universe.toggleCellAt(row+6, col+4)
		} else {
			// Toggle a single cell
			universe.toggleCellAt(row, col)
		}

		drawCanvas()
		return nil
	})

	playPauseButton := document.Call("getElementById", "play-pause")
	addEventListener("play-pause", "click", func(this js.Value, args []js.Value) interface{} {
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

	addEventListener("reset", "click", func(this js.Value, args []js.Value) interface{} {
		universe.reset()
		drawCanvas()
		return nil
	})

	addEventListener("randomize", "click", func(this js.Value, args []js.Value) interface{} {
		universe = NewUniverse()
		drawCanvas()
		return nil
	})

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

func addEventListener(elementID string, eventName string, callback func(this js.Value, args []js.Value) interface{}) {
	js.Global().
		Get("document").
		Call("getElementById", elementID).
		Call("addEventListener", eventName, js.FuncOf(callback))
}
