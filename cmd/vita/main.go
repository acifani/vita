package main

import (
	"flag"
	"fmt"

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

	for i := 0; i < *generations; i++ {
		fmt.Println(universe)
		fmt.Println()

		universe.Tick()
	}
}
