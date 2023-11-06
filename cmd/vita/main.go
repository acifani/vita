package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/acifani/vita/lib/game"
)

var (
	height      = flag.Int("height", 32, "height of the universe")
	width       = flag.Int("width", 32, "width of the universe")
	generations = flag.Int("gens", 3, "how many generations to run the universe")
	population  = flag.Int("pop", 45, "initial population percent of the universe")
	number      = flag.Int("n", 1, "number of universes to run in parallel")
)

func main() {
	flag.Parse()

	if *number > 1 {
		multi := createParallelUniverses()
		connectParallelUniverses(multi)
		runParallelUniverses(multi)

		return
	}

	runSingleUniverse()
}

func runSingleUniverse() {
	universe := game.NewUniverse(uint32(*height), uint32(*width))
	universe.Randomize(*population)

	for i := 0; i < *generations; i++ {
		fmt.Println(universe)
		fmt.Println()

		universe.Tick()
	}
}

func runParallelUniverses(multi []*game.ParallelUniverse) {
	for i := 0; i < *generations; i++ {
		for i, u := range multi {
			fmt.Println("Universe", i)
			fmt.Println(u)
			fmt.Println()
		}

		var wg sync.WaitGroup
		for _, u := range multi[:len(multi)] {
			callMultiTick(&wg, u)
		}
		wg.Wait()
	}
}

func createParallelUniverses() []*game.ParallelUniverse {
	multi := []*game.ParallelUniverse{}
	for row := 0; row < *number; row++ {
		for col := 0; col < *number; col++ {
			u := game.NewParallelUniverse(uint32(*height), uint32(*width))
			u.Randomize(*population)
			multi = append(multi, u)
		}
	}

	return multi
}

func connectParallelUniverses(multi []*game.ParallelUniverse) {
	i := 0
	for row := 0; row < *number; row++ {
		for col := 0; col < *number; col++ {
			switch {
			case col == *number-1 && row == *number-1:
				// do nothing
			case col == *number-1:
				multi[i].SetTopNeighbor(multi[*number+i])
			case row == *number-1:
				multi[i].SetRightNeighbor(multi[*number+1])
			default:
				multi[i].SetTopNeighbor(multi[*number+i])
				multi[i].SetRightNeighbor(multi[*number+1])
			}
		}
	}
}

func callMultiTick(wg *sync.WaitGroup, u *game.ParallelUniverse) {
	wg.Add(1)
	go func() {
		u.MultiTick()
		wg.Done()
	}()
}
