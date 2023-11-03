package main

import (
	"flag"

	"github.com/acifani/vita/lib/game"
)

func main() {
	height := flag.Int("height", 32, "height of the universe")
	width := flag.Int("width", 32, "width of the universe")
	generations := flag.Int("gens", 3, "how many generations to run the universe")
	population := flag.Int("pop", 45, "initial population percent of the universe")

	flag.Parse()

	universe := game.NewUniverse(uint32(*height), uint32(*width))
	universe.Randomize(*population)

	b := make([]byte, *height**width)

	for i := 0; i < *generations; i++ {
		universe.Read(b)
		output(b, *width)
		universe.Tick()
	}
}

func output(b []byte, width int) {
	for i := 0; i < len(b); i++ {
		if i%width == 0 {
			println()
		}
		if b[i] == game.Dead {
			print(".")
		} else {
			print("O")
		}
	}
	println()
	println()
}
