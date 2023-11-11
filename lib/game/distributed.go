package game

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
)

const (
	IDLength = 32
)

var NullID = strings.Repeat("0", IDLength)

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

func NewDistributedUniverse(id string, height, width uint32) *DistributedUniverse {
	d := &DistributedUniverse{
		ID:       id,
		TopID:    NullID,
		BottomID: NullID,
		LeftID:   NullID,
		RightID:  NullID,
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
				if d.TopID == NullID {
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
				if d.BottomID == NullID {
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
				if d.LeftID == NullID {
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
				if d.RightID == NullID {
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
	u := NewDistributedUniverse(id, d.height, d.width)
	return u
}

func (d *DistributedUniverse) Read(p []byte) (n int, err error) {
	if len(p) != IDLength*5+len(d.cells) {
		return 0, errInvalidLength
	}

	copy(p[:32], []byte(d.ID))
	copy(p[32:64], []byte(d.TopID))
	copy(p[64:96], []byte(d.BottomID))
	copy(p[96:128], []byte(d.LeftID))
	copy(p[128:160], []byte(d.RightID))
	copy(p[160:], d.cells)

	return len(p), nil
}

func (d *DistributedUniverse) Write(p []byte) (n int, err error) {
	if len(p) != IDLength*5+len(d.cells) {
		return 0, errInvalidLength
	}

	d.ID = string(p[:32])
	d.TopID = string(p[32:64])
	d.BottomID = string(p[64:96])
	d.LeftID = string(p[96:128])
	d.RightID = string(p[128:160])
	copy(d.cells, p[160:])

	return len(p), nil
}

// GenerateKey returns a string of length 32, since that
// is what you get from 16 bytes encoded as a hex string.
func GenerateKey() string {
	var result [16]byte
	rand.Read(result[:])
	return hex.EncodeToString(result[:])
}
