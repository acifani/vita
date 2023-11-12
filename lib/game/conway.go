package game

// ConwayRules implements the classic Game of Life rules.
func (u *Universe) ConwayRules(cell uint8, row, column uint32) uint8 {
	return RuleB3S23(cell, u.MooreNeighbors(row, column))
}

// ConwayRulesWrap implements the classic Game of Life rules but wraps the grid.
func (u *Universe) ConwayRulesWrap(cell uint8, row, column uint32) uint8 {
	return RuleB3S23(cell, u.MooreNeighborsWrap(row, column))
}

// RuleB3S23 implements the B3/S23 ruleset:
// https://www.conwaylife.com/wiki/Conway%27s_Game_of_Life#Rules
func RuleB3S23(cell uint8, liveNeighbors uint8) uint8 {
	switch {
	case cell == Dead && liveNeighbors == 3:
		// Birth - any dead cell with exactly three live neighbours
		// becomes a live cell, as if by reproduction.
		return Alive
	case cell == Alive && (liveNeighbors == 2 || liveNeighbors == 3):
		// Survival - any live cell with two or three live neighbours
		// lives on to the next generation.
		return Alive
	case cell == Alive && (liveNeighbors < 2 || liveNeighbors > 3):
		// Death - any live cell with less than two or more than three
		// live neighbours dies, as if by underpopulation/overpopulation.
		return Dead
	default:
		return cell
	}
}

// RuleB34S23 implements the B34/S23 ruleset:
// https://www.conwaylife.com/wiki/Conway%27s_Game_of_Life#Rules
func RuleB34S23(cell uint8, liveNeighbors uint8) uint8 {
	switch {
	case cell == Dead && (liveNeighbors == 3 || liveNeighbors == 4):
		// Birth - any dead cell with three or four live neighbours 
		// becomes a live cell as if by reproduction.
		return Alive
	case cell == Alive && (liveNeighbors == 2 || liveNeighbors == 3):
		// Survival - any live cell with two or three live neighbours
		// lives on to the next generation.
		return Alive
	case cell == Alive && (liveNeighbors < 2 || liveNeighbors > 3):
		// Death - any live cell with less than two or more than three
		// live neighbours dies, as if by underpopulation/overpopulation.
		return Dead
	default:
		return cell
	}
}

// MooreNeighbors returns the number of alive neighbors for a given cell.
// It uses the Moore neighborhood, which includes the eight cells surrounding
// the given cell.
func (u *Universe) MooreNeighbors(row, column uint32) uint8 {
	count := uint8(0)

	r, c := int32(row), int32(column)
	for _, neighborRow := range []int32{r - 1, r, r + 1} {
		for _, neighborColumn := range []int32{c - 1, c, c + 1} {
			switch {
			// Skip checking the cell itself
			case neighborRow == r && neighborColumn == c:
			// Ignore if pixel is out of bounds
			case neighborRow < 0:
			case neighborColumn < 0:
			case neighborRow > int32(u.height)-1:
			case neighborColumn > int32(u.width)-1:
			// otherwise, check neighbor
			default:
				neighborIdx := u.GetIndex(uint32(neighborRow), uint32(neighborColumn))
				if u.Cell(neighborIdx) != Dead {
					count++
				}
			}
		}
	}

	return count
}

// MooreNeighborsWrap returns the number of alive neighbors for a given cell.
// It uses the Moore neighborhood, which includes the eight cells surrounding
// the given cell, but wraps to the other side if it would be off the grid.
func (u *Universe) MooreNeighborsWrap(row, column uint32) uint8 {
	count := uint8(0)

	r, c := int32(row), int32(column)
	for _, neighborRow := range []int32{r - 1, r, r + 1} {
		for _, neighborColumn := range []int32{c - 1, c, c + 1} {
			switch {
			// Skip checking the cell itself
			case neighborRow == r && neighborColumn == c:
				continue
			// Wrap if pixel is out of bounds
			case neighborRow < 0:
				neighborRow = int32(u.height) - 1
			case neighborColumn < 0:
				neighborColumn = int32(u.width) - 1
			case neighborRow > int32(u.height)-1:
				neighborRow = 0
			case neighborColumn > int32(u.width)-1:
				neighborColumn = 0
			}
			// Now, check neighbor
			neighborIdx := u.GetIndex(uint32(neighborRow), uint32(neighborColumn))
			if u.Cell(neighborIdx) != Dead {
				count++
			}
		}
	}

	return count
}
