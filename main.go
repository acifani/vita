package main

import (
	"math"
	"strconv"
	"syscall/js"
)

const (
	dead  = iota
	alive = iota
)

const (
	toggleAction = iota
	gliderAction = iota
	pulsarAction = iota
)

const (
	cellSize   = 10
	borderSize = 1
)

var (
	universe       *Universe
	ctx            js.Value
	lastTick       float64
	animationID    = -1
	clickAction    = toggleAction
	livePopulation = 50
	renderingSpeed = 50
)

func main() {
	done := make(chan bool)

	universe = NewUniverse(livePopulation)
	window := js.Global()
	document := window.Get("document")

	canvas := setupCanvas()
	ctx = canvas.Call("getContext", "2d")

	gps := document.Call("getElementById", "gps")
	ticks := float64(0)
	renderingLoops := 0

	// Rendering loop
	var draw js.Func
	draw = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		renderingLoops = renderingLoops + 1
		if renderingLoops > (5 - renderingSpeed) {
			universe.tick()
			renderingLoops = 0

			ticks = ticks + 1

			if ticks >= 20 {
				now := window.Get("performance").Call("now").Float()
				elapsedMs := now - lastTick
				lastTick = now
				// Number of frames divided by the seconds that have passed
				generationsPerSecond := ticks / (elapsedMs / 1000)
				gps.Set("innerText", int(generationsPerSecond))
				ticks = 0
			}
		}

		drawCanvas()
		animationID = window.Call("requestAnimationFrame", draw).Int()
		return nil
	})
	defer draw.Release()

	// Add event listeners

	addEventListener("live-population", "change", func(this js.Value, args []js.Value) interface{} {
		newValue := args[0].Get("target").Get("value").String()
		livePopulation, _ = strconv.Atoi(newValue)
		universe = NewUniverse(livePopulation)
		return nil
	})

	addEventListener("rendering-speed", "change", func(this js.Value, args []js.Value) interface{} {
		newValue := args[0].Get("target").Get("value").String()
		renderingSpeed, _ = strconv.Atoi(newValue)
		drawCanvas()
		return nil
	})

	addEventListener("canvas", "click", func(this js.Value, args []js.Value) interface{} {
		boundingRect := canvas.Call("getBoundingClientRect")
		widthScale := canvas.Get("width").Int() / boundingRect.Get("width").Int()
		heightScale := canvas.Get("height").Int() / boundingRect.Get("height").Int()
		canvasX := (args[0].Get("clientX").Int() - boundingRect.Get("left").Int()) * widthScale
		canvasY := (args[0].Get("clientY").Int() - boundingRect.Get("top").Int()) * heightScale
		row := uint32(math.Floor(float64(canvasY) / (cellSize + borderSize)))
		col := uint32(math.Floor(float64(canvasX) / (cellSize + borderSize)))

		switch clickAction {
		case gliderAction:
			figure := glider()
			universe.setRectangle(row-figure.deltaX, col-figure.deltaY, figure.values)
		case pulsarAction:
			figure := pulsar()
			universe.setRectangle(row-figure.deltaX, col-figure.deltaY, figure.values)
		default:
			universe.toggleCellAt(row, col)
		}

		drawCanvas()
		return nil
	})

	addEventListener("toggle", "click", func(this js.Value, args []js.Value) interface{} {
		clickAction = toggleAction
		return nil
	})

	addEventListener("glider", "click", func(this js.Value, args []js.Value) interface{} {
		clickAction = gliderAction
		return nil
	})

	addEventListener("pulsar", "click", func(this js.Value, args []js.Value) interface{} {
		clickAction = pulsarAction
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
		universe = NewUniverse(livePopulation)
		drawCanvas()
		return nil
	})

	// Start rendering
	lastTick = window.Get("performance").Call("now").Float()
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
	ctx.Set("strokeStyle", "#e0e1e4")

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
	ctx.Set("fillStyle", "#3c4257")
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
	ctx.Set("fillStyle", "#fff")
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
