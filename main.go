package main

import (
	"fmt"
	"math/rand"
)

const (
	dead  = iota
	alive = iota
)

func main() {
	universe := NewUniverse()
	fmt.Print(universe.toString())
}

type Universe struct {
	height uint32
	width  uint32
	cells  []uint8
}

func NewUniverse() *Universe {
	width := uint32(64)
	height := uint32(64)
	cells := make([]uint8, width*height)

	for i := range cells {
		if rand.Intn(100) < 50 {
			cells[i] = alive
		} else {
			cells[i] = dead
		}
	}

	return &Universe{
		height: height,
		width:  width,
		cells:  cells,
	}
}

func (u *Universe) toString() string {
	universeString := ""
	for i, cell := range u.cells {
		if cell == alive {
			universeString += "◼"
		} else {
			universeString += "◻"
		}

		if i%int(u.width) == 0 {
			universeString += "\n"
		}
	}

	return universeString
}
