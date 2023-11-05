package game

// ConwayRules implements the classic Game of Life rules.
func (u *Universe) ConwayRules(cell uint8, row, column uint32) uint8 {
	return RuleB3S23(cell, u.MooreNeighbors(row, column))
}

// RuleB3S23 implements the B3/S23 ruleset:
// https://www.conwaylife.com/wiki/Conway%27s_Game_of_Life#Rules
func RuleB3S23(cell uint8, liveNeighbors uint8) uint8 {
	switch {
	case cell == Alive && liveNeighbors < 2:
		// 1. Any live cell with fewer than two live neighbours dies, as if by underpopulation.
		return Dead
	case cell == Alive && (liveNeighbors == 2 || liveNeighbors == 3):
		// 2. Any live cell with two or three live neighbours lives on to the next generation.
		return Alive
	case cell == Alive && liveNeighbors > 3:
		// 3. Any live cell with more than three live neighbours dies, as if by overpopulation.
		return Dead
	case cell == Dead && liveNeighbors == 3:
		// 4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
		return Alive
	default:
		return cell
	}
}

// MooreNeighbors returns the number of alive neighbors for a given cell.
// It uses the Moore neighborhood, which includes the eight cells surrounding
// the given cell.
func (u *Universe) MooreNeighbors(row, column uint32) uint8 {
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

			// If we are in a multiverse,
			// and we are at column 0,
			// and we would check neighbors at column-1,
			// then we want to check the information given by the other universe instead
			if u.multiverse && column == 0 && colDiff == u.width-1 {
				if u.multiverseColumn[neighborRow] == Alive {
					count++
					continue
				}
			}

			if u.Cell(neighborIdx) == Alive {
				count++
			}
		}
	}

	return count
}
