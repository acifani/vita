package game

import (
	"strings"
)

const (
	Dead = iota
	Alive
)

type Universe struct {
	height     uint32
	width      uint32
	cells      []uint8
	newCells   []uint8
	stable     bool
	Generation uint32

	Rules func(cell uint8, row, column uint32) uint8
}

func NewUniverse(height, width uint32) *Universe {
	cells := make([]uint8, height*width)
	newCells := make([]uint8, height*width)

	u := &Universe{
		height:   height,
		width:    width,
		cells:    cells,
		newCells: newCells,
	}
	u.Rules = u.ConwayRules
	return u
}

func (u *Universe) Height() uint32 {
	return u.height
}

func (u *Universe) Width() uint32 {
	return u.width
}

func (u *Universe) Size() int {
	return len(u.cells)
}

func (u *Universe) Dead() bool {
	for i := range u.cells {
		if u.cells[i] == Alive {
			return false
		}
	}

	return true
}

// Stable returns true if the universe has reached a stable state.
// A stable state is one in which no cells have changed between ticks.
// See https://conwaylife.com/wiki/Still_life for more information.
func (u *Universe) Stable() bool {
	return u.stable
}

func (u *Universe) Cell(idx uint32) uint8 {
	return u.cells[idx]
}

func (u *Universe) GetIndex(row, column uint32) uint32 {
	return row*u.width + column
}

func (u *Universe) Tick() {
	stable := true
	for row := uint32(0); row < u.height; row++ {
		for column := uint32(0); column < u.width; column++ {
			cellIndex := u.GetIndex(row, column)
			cell := u.cells[cellIndex]

			u.newCells[cellIndex] = u.Rules(cell, row, column)
			if u.newCells[cellIndex] != cell {
				stable = false
			}
		}
	}

	u.stable = stable
	u.Generation++
	copy(u.cells, u.newCells)
}

func (u *Universe) Reset() {
	for i := range u.cells {
		u.cells[i] = Dead
	}
	u.Generation = 0
}

func (u *Universe) Randomize(livePopulation int) {
	for i := range u.cells {
		if randomNumber() < livePopulation {
			u.cells[i] = Alive
		} else {
			u.cells[i] = Dead
		}
	}
}

func (u *Universe) ToggleCellAt(row, column uint32) {
	idx := u.GetIndex(row, column)
	if u.cells[idx] == Alive {
		u.cells[idx] = Dead
	} else {
		u.cells[idx] = Alive
	}
}

func (u *Universe) SetRectangle(startingRow, startingColumn uint32, values [][]uint8) {
	for i, row := range values {
		for j, value := range row {
			idx := u.GetIndex(startingRow+uint32(i), startingColumn+uint32(j))
			u.cells[idx] = value
		}
	}
}

func (u *Universe) Read(p []byte) (n int, err error) {
	if len(p) != u.Size() {
		return 0, errInvalidLength
	}

	copy(p, u.cells)
	return len(p), nil
}

func (u *Universe) Write(p []byte) (n int, err error) {
	if len(p) != u.Size() {
		return 0, errInvalidLength
	}

	copy(u.cells, p)
	return len(p), nil
}

func (u *Universe) String() string {
	builder := strings.Builder{}
	for i := 0; i < len(u.cells); i++ {
		if i%int(u.width) == 0 && i != 0 {
			builder.WriteString("\n")
		}
		if u.cells[i] == Dead {
			builder.WriteString(".")
		} else {
			builder.WriteString("O")
		}
	}
	builder.WriteString("\n")

	return builder.String()
}
