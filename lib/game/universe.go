package game

import (
	"math/rand"
)

const (
	Dead = iota
	Alive
)

type Universe struct {
	height   uint32
	width    uint32
	cells    []uint8
	newCells []uint8

	Rule func(state uint8, liveNeighbors uint8) uint8
}

func NewUniverse(height, width uint32) *Universe {
	cells := make([]uint8, height*width)
	newCells := make([]uint8, height*width)

	return &Universe{
		height:   height,
		width:    width,
		cells:    cells,
		newCells: newCells,
		Rule:     rule,
	}
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

func (u *Universe) Cell(idx uint32) uint8 {
	return u.cells[idx]
}

func (u *Universe) GetIndex(row, column uint32) uint32 {
	return row*u.width + column
}

func (u *Universe) AliveNeighbors(row, column uint32) uint8 {
	count := uint8(0)

	// We use height/width and modulos to avoid manually handling edge cases
	// (literally: cases at the edge of the universe, e.g. cells at row/col 0).
	for _, rowDiff := range []uint32{u.height - 1, 0, 1} {
		for _, colDiff := range []uint32{u.width - 1, 0, 1} {
			if rowDiff == 0 && colDiff == 0 {
				// Skip checking the cell itself
				continue
			}

			neighborRow := (row + rowDiff) % u.height
			neighborColumn := (column + colDiff) % u.width
			neighborIdx := u.GetIndex(neighborRow, neighborColumn)
			count += u.cells[neighborIdx]
		}
	}

	return count
}

func (u *Universe) Tick() {
	for row := uint32(0); row < u.height; row++ {
		for column := uint32(0); column < u.width; column++ {
			cellIndex := u.GetIndex(row, column)
			cell := u.cells[cellIndex]
			liveNeighbors := u.AliveNeighbors(row, column)

			u.newCells[cellIndex] = u.Rule(cell, liveNeighbors)
		}
	}

	copy(u.cells, u.newCells)
}

func (u *Universe) Reset() {
	for i := range u.cells {
		u.cells[i] = Dead
	}
}

func (u *Universe) Randomize(livePopulation int) {
	for i := range u.cells {
		if rand.Intn(100) < livePopulation {
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

func rule(state uint8, liveNeighbors uint8) uint8 {
	switch {
	case state == Alive && liveNeighbors < 2:
		// 1. Any live cell with fewer than two live neighbours dies, as if by underpopulation.
		return Dead
	case state == Alive && (liveNeighbors == 2 || liveNeighbors == 3):
		// 2. Any live cell with two or three live neighbours lives on to the next generation.
		return Alive
	case state == Alive && liveNeighbors > 3:
		// 3. Any live cell with more than three live neighbours dies, as if by overpopulation.
		return Dead
	case state == Dead && liveNeighbors == 3:
		// 4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
		return Alive
	default:
		return state
	}
}
