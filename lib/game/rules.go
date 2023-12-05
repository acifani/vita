package game

// ConwayRules implements the classic Game of Life rules.
func (u *Universe) ConwayRules(cell uint8, row, column uint32) uint8 {
	return RuleB3S23(cell, u.MooreNeighbors(row, column))
}

// ConwayRulesWrap implements the classic Game of Life rules but wraps the grid.
func (u *Universe) ConwayRulesWrap(cell uint8, row, column uint32) uint8 {
	return RuleB3S23(cell, u.MooreNeighborsWrap(row, column))
}

// RuleB3S23 implements the B3/S23 rules:
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

// SeedsRules implements the Seeds rules.
// See https://conwaylife.com/wiki/OCA:Seeds
func (u *Universe) SeedsRules(cell uint8, row, column uint32) uint8 {
	return RuleB2S(cell, u.MooreNeighbors(row, column))
}

// RuleB2S implements the B2S rules:
// https://conwaylife.com/wiki/OCA:Seeds
func RuleB2S(cell uint8, liveNeighbors uint8) uint8 {
	switch {
	case cell == Dead && (liveNeighbors == 2):
		// Birth - any dead cell with two live neighbours
		// becomes a live cell as if by reproduction.
		return Alive
	case cell == Alive:
		// Death - any live cell dies.
		return Dead
	default:
		return cell
	}
}

// DayAndNightRules implements the Day And Night rules.
// See https://conwaylife.com/wiki/OCA:Day_%26_Night
func (u *Universe) DayAndNightRules(cell uint8, row, column uint32) uint8 {
	return RuleB3678S34678(cell, u.MooreNeighbors(row, column))
}

// RuleB3678S34678 implements the B3678/S34678 rules:
// https://conwaylife.com/wiki/OCA:Day_%26_Night
func RuleB3678S34678(cell uint8, liveNeighbors uint8) uint8 {
	switch {
	case cell == Dead && (liveNeighbors == 3 || liveNeighbors == 6 ||
		liveNeighbors == 7 || liveNeighbors == 8):
		// Birth - any dead cell with 3, 6, 7, or 8 live neighbours
		// becomes a live cell as if by reproduction.
		return Alive
	case cell == Alive && (liveNeighbors == 3 || liveNeighbors == 4 ||
		liveNeighbors == 6 || liveNeighbors == 7 || liveNeighbors == 8):
		// Survival - any live cell with 3, 4, 6, 7, or 8 live neighbours
		// lives on to the next generation.
		return Alive
	case cell == Alive && (liveNeighbors < 3 || liveNeighbors == 5 || liveNeighbors > 8):
		// Death - any live cell with less than three, five, or more than eight
		// live neighbours dies, as if by underpopulation/overpopulation.
		return Dead
	default:
		return cell
	}
}

// WolframRule30 implements Rule 30 from Stephen Wolfram's "A New Kind of Science".
// See https://en.wikipedia.org/wiki/Rule_30
func (u *Universe) WolframRule30(cell uint8, row, column uint32) uint8 {
	prev, next := u.OneDimensionalNeighbors(row, column)

	switch cell {
	case Alive:
		switch {
		case prev == Alive && next == Alive:
			return Dead
		case prev == Dead && next == Dead:
			return Alive
		case next == Alive:
			return Alive
		default:
			return Dead
		}
	case Dead:
		switch {
		case prev == Alive && next == Alive:
			return Dead
		case prev == Alive || next == Alive:
			return Alive
		default:
			return Dead
		}
	default:
		return cell
	}
}

// WolframRule90 implements Rule 90 from Stephen Wolfram's "A New Kind of Science".
// See https://en.wikipedia.org/wiki/Rule_90
func (u *Universe) WolframRule90(cell uint8, row, column uint32) uint8 {
	prev, next := u.OneDimensionalNeighbors(row, column)

	switch cell {
	case Alive:
		switch {
		case prev == Alive && next == Alive:
			return Dead
		case prev == Alive || next == Alive:
			return Alive
		case prev == Dead && next == Dead:
			return Dead
		default:
			return Alive
		}
	case Dead:
		switch {
		case prev == Alive && next == Alive:
			return Dead
		case prev == Alive || next == Alive:
			return Alive
		default:
			return Dead
		}
	default:
		return cell
	}
}

// WolframRule110 implements Rule 110 from Stephen Wolfram's "A New Kind of Science".
// See https://en.wikipedia.org/wiki/Rule_110
func (u *Universe) WolframRule110(cell uint8, row, column uint32) uint8 {
	prev, next := u.OneDimensionalNeighbors(row, column)

	switch cell {
	case Alive:
		switch {
		case prev == Alive && next == Alive:
			return Dead
		default:
			return Alive
		}
	case Dead:
		switch {
		case prev == Alive && next == Alive:
			return Alive
		case next == Alive:
			return Alive
		default:
			return Dead
		}
	default:
		return cell
	}
}

// WolframRule184 implements Rule 184 from Stephen Wolfram's "A New Kind of Science".
// See https://en.wikipedia.org/wiki/Rule_184
func (u *Universe) WolframRule184(cell uint8, row, column uint32) uint8 {
	prev, next := u.OneDimensionalNeighbors(row, column)

	switch cell {
	case Alive:
		switch {
		case next == Alive:
			return Alive
		default:
			return Dead
		}
	case Dead:
		switch {
		case prev == Alive:
			return Alive
		default:
			return Dead
		}
	default:
		return cell
	}
}

func (u *Universe) OneDimensionalNeighbors(row, column uint32) (uint8, uint8) {
	var prev, next uint8
	switch {
	case column == 0:
		prev = Dead
		next = u.Cell(u.GetIndex(row, column+1))
	case column >= u.width-1:
		prev = u.Cell(u.GetIndex(row, column-1))
		next = Dead
	default:
		prev = u.Cell(u.GetIndex(row, column-1))
		next = u.Cell(u.GetIndex(row, column+1))
	}

	return prev, next
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
