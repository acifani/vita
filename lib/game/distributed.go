package game

type DistributedUniverse struct {
	*Universe

	ID             string
	TopID          string
	TopNeighbor    *DistributedUniverse
	BottomID       string
	BottomNeighbor *DistributedUniverse
	LeftID         string
	LeftNeighbor   *DistributedUniverse
	RightID        string
	RightNeighbor  *DistributedUniverse

	// set the GetNeighbor function with implementation for
	// the distributed environment.
	GetNeighbor func(id string) *DistributedUniverse
}

func NewDistributedUniverse(height, width uint32) *DistributedUniverse {
	d := &DistributedUniverse{
		Universe: NewUniverse(height, width),
	}
	d.Rules = d.rules
	d.GetNeighbor = d.getEmptyUniverse
	return d
}

func (d *DistributedUniverse) rules(cell uint8, row, column uint32) uint8 {
	return RuleB3S23(cell, d.Neighbors(row, column))
}

// Neighbors returns the number of alive neighbors for a given cell.
// It uses the Moore neighborhood, which includes the eight cells surrounding
// the given cell, but also extended to adjoining neighbor Universes.
func (d *DistributedUniverse) Neighbors(row, column uint32) uint8 {
	count := uint8(0)
	r, c := int32(row), int32(column)
	for _, neighborRow := range []int32{r - 1, r, r + 1} {
		for _, neighborColumn := range []int32{c - 1, c, c + 1} {
			switch {
			case neighborRow == r && neighborColumn == c:
				// Skip checking the cell itself
				continue
			case neighborRow < 0 && neighborColumn < 0:
				// TODO: check the universe above and to the left for 1 pixel?
			case neighborRow < 0 && neighborColumn >= int32(d.width):
				// TODO: check the universe above and to the right for 1 pixel?
			case neighborRow >= int32(d.height) && neighborColumn < 0:
				// TODO: check the universe below and to the left for 1 pixel?
			case neighborRow >= int32(d.height) && neighborColumn >= int32(d.width):
				// check the universe below and to the right for one pixel?
			case neighborRow < 0:
				// check the universe above
				if d.TopID == "" {
					continue
				}
				if d.TopNeighbor == nil {
					d.TopNeighbor = d.GetNeighbor(d.TopID)
				}

				neighborIdx := d.GetIndex(d.height-1, uint32(neighborColumn))
				if d.TopNeighbor.Cell(neighborIdx) != Dead {
					count++
				}
			case neighborRow >= int32(d.height):
				// check the universe below
				if d.BottomID == "" {
					continue
				}
				if d.BottomNeighbor == nil {
					d.BottomNeighbor = d.GetNeighbor(d.BottomID)
				}

				neighborIdx := d.GetIndex(0, uint32(neighborColumn))
				if d.BottomNeighbor.Cell(neighborIdx) != Dead {
					count++
				}
			case neighborColumn < 0:
				// check the universe to the left
				if d.LeftID == "" {
					continue
				}
				if d.LeftNeighbor == nil {
					d.LeftNeighbor = d.GetNeighbor(d.LeftID)
				}

				neighborIdx := d.GetIndex(uint32(neighborRow), d.width-1)
				if d.LeftNeighbor.Cell(neighborIdx) != Dead {
					count++
				}
			case neighborColumn >= int32(d.width):
				// check the universe to the right
				if d.RightID == "" {
					continue
				}
				if d.RightNeighbor == nil {
					d.RightNeighbor = d.GetNeighbor(d.RightID)
				}

				neighborIdx := d.GetIndex(uint32(neighborRow), 0)
				if d.RightNeighbor.Cell(neighborIdx) != Dead {
					count++
				}
			default:
				// check the current universe
				neighborIdx := d.GetIndex(uint32(neighborRow), uint32(neighborColumn))
				if d.Cell(neighborIdx) != Dead {
					count++
				}
			}
		}
	}

	return count
}

func (d *DistributedUniverse) SetTopNeighbor(n *DistributedUniverse) error {
	if d.ID == n.ID {
		return errInvalidID
	}

	d.TopID = n.ID
	return nil
}

func (d *DistributedUniverse) SetBottomNeighbor(n *DistributedUniverse) error {
	if d.ID == n.ID {
		return errInvalidID
	}

	d.BottomID = n.ID
	return nil
}

func (d *DistributedUniverse) SetLeftNeighbor(n *DistributedUniverse) error {
	if d.ID == n.ID {
		return errInvalidID
	}

	d.LeftID = n.ID
	return nil
}

func (d *DistributedUniverse) SetRightNeighbor(n *DistributedUniverse) error {
	if d.ID == n.ID {
		return errInvalidID
	}

	d.RightID = n.ID
	return nil
}

func (d *DistributedUniverse) getEmptyUniverse(id string) *DistributedUniverse {
	u := NewDistributedUniverse(d.height, d.width)
	u.ID = id
	return u
}
