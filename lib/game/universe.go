package game

import "math/rand"

const (
	Dead = iota
	Alive
)

type Universe struct {
	height   uint32
	width    uint32
	cells    []uint8
	newCells []uint8
}

func NewUniverse(livePopulation int) *Universe {
	width := uint32(64)
	height := uint32(64)
	cells := make([]uint8, width*height)
	newCells := make([]uint8, width*height)

	for i := range cells {
		if rand.Intn(100) < livePopulation {
			cells[i] = Alive
		} else {
			cells[i] = Dead
		}
	}

	return &Universe{
		height:   height,
		width:    width,
		cells:    cells,
		newCells: newCells,
	}
}

func (u *Universe) Height() uint32 {
	return u.height
}

func (u *Universe) Width() uint32 {
	return u.width
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
	for row := uint32(0); row < u.width; row++ {
		for column := uint32(0); column < u.height; column++ {
			cellIndex := u.GetIndex(row, column)
			cell := u.cells[cellIndex]
			liveNeighbors := u.AliveNeighbors(row, column)

			switch {
			case cell == Alive && liveNeighbors < 2:
				// 1. Any live cell with fewer than two live neighbours dies, as if by underpopulation.
				u.newCells[cellIndex] = Dead
			case cell == Alive && (liveNeighbors == 2 || liveNeighbors == 3):
				// 2. Any live cell with two or three live neighbours lives on to the next generation.
				u.newCells[cellIndex] = Alive
			case cell == Alive && liveNeighbors > 3:
				// 3. Any live cell with more than three live neighbours dies, as if by overpopulation.
				u.newCells[cellIndex] = Dead
			case cell == Dead && liveNeighbors == 3:
				// 4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
				u.newCells[cellIndex] = Alive
			default:
				u.newCells[cellIndex] = cell
			}
		}
	}

	copy(u.cells, u.newCells)
}

func (u *Universe) Reset() {
	for row := uint32(0); row < u.width; row++ {
		for column := uint32(0); column < u.height; column++ {
			u.cells[u.GetIndex(row, column)] = Dead
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
