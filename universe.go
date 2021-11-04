package main

import "math/rand"

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

func (u *Universe) getIndex(row, column uint32) uint32 {
	return row*u.width + column
}

func (u *Universe) aliveNeighbors(row, column uint32) uint8 {
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
			neighborIdx := u.getIndex(neighborRow, neighborColumn)
			count += u.cells[neighborIdx]
		}
	}

	return count
}

func (u *Universe) tick() {
	newCells := make([]uint8, u.height*u.width)

	for row := uint32(0); row < u.width; row++ {
		for column := uint32(0); column < u.height; column++ {
			cellIndex := u.getIndex(row, column)
			cell := u.cells[cellIndex]
			liveNeighbors := u.aliveNeighbors(row, column)

			if cell == alive && liveNeighbors < 2 {
				// 1. Any live cell with fewer than two live neighbours dies, as if by underpopulation.
				newCells[cellIndex] = dead
			} else if cell == alive && (liveNeighbors == 2 || liveNeighbors == 3) {
				// 2. Any live cell with two or three live neighbours lives on to the next generation.
				newCells[cellIndex] = alive
			} else if cell == alive && liveNeighbors > 3 {
				// 3. Any live cell with more than three live neighbours dies, as if by overpopulation.
				newCells[cellIndex] = dead
			} else if cell == dead && liveNeighbors == 3 {
				// 4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
				newCells[cellIndex] = alive
			} else {
				newCells[cellIndex] = cell
			}
		}
	}

	u.cells = newCells
}

func (u *Universe) toString() string {
	universeString := ""
	for i, cell := range u.cells {
		if cell == alive {
			universeString += "◼"
		} else {
			universeString += "◻"
		}

		if (i+1)%int(u.width) == 0 {
			universeString += "\n"
		}
	}

	return universeString
}

func (u *Universe) reset() {
	u.cells = make([]uint8, u.width*u.height)
}
