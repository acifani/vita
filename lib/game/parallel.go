package game

import (
	"sync"
)

type NeighborData struct {
	Generation uint32
	Cells      []uint8
}

type NeighborConnection struct {
	SendCh    chan NeighborData
	ReceiveCh chan NeighborData
	Data      NeighborData
}

func NewNeighborConnection() *NeighborConnection {
	return &NeighborConnection{
		SendCh: make(chan NeighborData, 1),
	}
}

type ParallelUniverse struct {
	*Universe

	TopNeighbor    *NeighborConnection
	BottomNeighbor *NeighborConnection
	LeftNeighbor   *NeighborConnection
	RightNeighbor  *NeighborConnection

	sendBuffer []uint8
}

func NewParallelUniverse(height, width uint32) *ParallelUniverse {
	p := &ParallelUniverse{
		Universe:   NewUniverse(height, width),
		sendBuffer: make([]uint8, height*width),
	}
	p.Rules = p.rules
	return p
}

func (p *ParallelUniverse) rules(cell uint8, row, column uint32) uint8 {
	return RuleB3S23(cell, p.Neighbors(row, column))
}

// The ParallelUniverse Neighbors function returns the number of alive neighbors
// for a given cell. It uses the Moore neighborhood, which includes the eight cells
// surrounding the given cell, but also extended to adjoining neighbor Universes.
func (p *ParallelUniverse) Neighbors(row, column uint32) uint8 {
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
			case neighborRow < 0 && neighborColumn >= int32(p.width):
				// TODO: check the universe above and to the right for 1 pixel?
			case neighborRow >= int32(p.height) && neighborColumn < 0:
				// TODO: check the universe below and to the left for 1 pixel?
			case neighborRow >= int32(p.height) && neighborColumn >= int32(p.width):
				// check the universe below and to the right for one pixel?
			case neighborRow < 0:
				// check the universe above
				if p.TopNeighbor == nil {
					continue
				}

				neighborIdx := p.GetIndex(p.height-1, uint32(neighborColumn))
				if p.TopNeighbor.Data.Cells[neighborIdx] != Dead {
					count++
				}
			case neighborRow >= int32(p.height):
				// check the universe below
				if p.BottomNeighbor == nil {
					continue
				}

				neighborIdx := p.GetIndex(0, uint32(neighborColumn))
				if p.BottomNeighbor.Data.Cells[neighborIdx] != Dead {
					count++
				}
			case neighborColumn < 0:
				// check the universe to the left
				if p.LeftNeighbor == nil {
					continue
				}

				neighborIdx := p.GetIndex(uint32(neighborRow), p.width-1)
				if p.LeftNeighbor.Data.Cells[neighborIdx] != Dead {
					count++
				}
			case neighborColumn >= int32(p.width):
				// check the universe to the right
				if p.RightNeighbor == nil {
					continue
				}

				neighborIdx := p.GetIndex(uint32(neighborRow), 0)
				if p.RightNeighbor.Data.Cells[neighborIdx] != Dead {
					count++
				}
			default:
				// check the current universe
				neighborIdx := p.GetIndex(uint32(neighborRow), uint32(neighborColumn))
				if p.Cell(neighborIdx) != Dead {
					count++
				}
			}
		}
	}

	return count
}

func (p *ParallelUniverse) MultiTick() {
	p.SendDataToNeighbors()
	p.WaitForNeighborsData()
	p.Tick()
}

// WaitForNeighborsData waits for all neighbors to send their data.
func (p *ParallelUniverse) WaitForNeighborsData() {
	var wg sync.WaitGroup
	if p.TopNeighbor != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			p.TopNeighbor.Data = <-p.TopNeighbor.ReceiveCh
		}()
	}
	if p.BottomNeighbor != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			p.BottomNeighbor.Data = <-p.BottomNeighbor.ReceiveCh
		}()
	}
	if p.LeftNeighbor != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			p.LeftNeighbor.Data = <-p.LeftNeighbor.ReceiveCh
		}()
	}
	if p.RightNeighbor != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			p.RightNeighbor.Data = <-p.RightNeighbor.ReceiveCh
		}()
	}

	wg.Wait()
}

func (p *ParallelUniverse) SendDataToNeighbors() {
	var wg sync.WaitGroup
	data := NeighborData{
		Generation: p.Generation,
		Cells:      p.sendBuffer,
	}
	copy(p.sendBuffer, p.cells)

	if p.TopNeighbor != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			p.TopNeighbor.SendCh <- data
		}()
	}
	if p.BottomNeighbor != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			p.BottomNeighbor.SendCh <- data
		}()
	}
	if p.LeftNeighbor != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			p.LeftNeighbor.SendCh <- data
		}()
	}
	if p.RightNeighbor != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			p.RightNeighbor.SendCh <- data
		}()
	}

	wg.Wait()
}

func (p *ParallelUniverse) SetTopNeighbor(n *ParallelUniverse) {
	if p.TopNeighbor == nil {
		p.TopNeighbor = NewNeighborConnection()
	}
	if n.BottomNeighbor == nil {
		n.BottomNeighbor = NewNeighborConnection()
	}

	p.TopNeighbor.ReceiveCh = n.BottomNeighbor.SendCh
}

func (p *ParallelUniverse) SetBottomNeighbor(n *ParallelUniverse) {
	if p.BottomNeighbor == nil {
		p.BottomNeighbor = NewNeighborConnection()
	}
	if n.TopNeighbor == nil {
		n.TopNeighbor = NewNeighborConnection()
	}

	p.BottomNeighbor.ReceiveCh = n.TopNeighbor.SendCh
}

func (p *ParallelUniverse) SetLeftNeighbor(n *ParallelUniverse) {
	if p.LeftNeighbor == nil {
		p.LeftNeighbor = NewNeighborConnection()
	}
	if n.RightNeighbor == nil {
		n.RightNeighbor = NewNeighborConnection()
	}

	p.LeftNeighbor.ReceiveCh = n.RightNeighbor.SendCh
}

func (p *ParallelUniverse) SetRightNeighbor(n *ParallelUniverse) {
	if p.RightNeighbor == nil {
		p.RightNeighbor = NewNeighborConnection()
	}
	if n.LeftNeighbor == nil {
		n.LeftNeighbor = NewNeighborConnection()
	}

	p.RightNeighbor.ReceiveCh = n.LeftNeighbor.SendCh
}
