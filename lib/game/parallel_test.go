package game

import (
	"testing"
)

func TestParallelUniverse(t *testing.T) {
	t.Run("NewParallelUniverse", func(t *testing.T) {
		u := NewParallelUniverse(24, 32)

		if u.Height() != 24 {
			t.Errorf("Expected height to be 24, got %d", u.Height())
		}

		if u.Width() != 32 {
			t.Errorf("Expected width to be 32, got %d", u.Width())
		}

		var i uint32
		for i = 0; i < u.Height()*u.Width(); i++ {
			if u.Cell(i) != Dead {
				t.Errorf("Expected cell %d to be dead, got %d", i, u.Cell(uint32(i)))
			}
		}
	})

	t.Run("TickWithoutNeighbors", func(t *testing.T) {
		u := NewParallelUniverse(24, 32)
		u.cells[u.GetIndex(10, 12)] = Alive
		u.cells[u.GetIndex(11, 12)] = Alive
		u.cells[u.GetIndex(12, 12)] = Alive
		u.cells[u.GetIndex(11, 13)] = Alive

		u.Tick()

		if u.Stable() {
			t.Errorf("Expected universe to not be stable")
		}

		if u.ParallelNeighbors(0, 0) != 0 {
			t.Errorf("Expected cell %d to have 0 alive neighbors, got %d", u.GetIndex(0, 0), u.ParallelNeighbors(0, 0))
		}

		if u.ParallelNeighbors(11, 12) != 6 {
			t.Errorf("Expected cell %d to have 6 alive neighbors, got %d", u.GetIndex(11, 12), u.ParallelNeighbors(11, 12))
		}

		u.Tick()

		if u.Stable() {
			t.Errorf("Expected universe to not be stable")
		}

		if u.Cell(u.GetIndex(11, 12)) != Dead {
			t.Errorf("Expected cell %d to be dead, got %d", u.GetIndex(11, 12), u.Cell(u.GetIndex(11, 12)))
		}

		if u.ParallelNeighbors(11, 12) != 5 {
			t.Errorf("Expected cell %d to have 5 alive neighbors, got %d", u.GetIndex(11, 12), u.ParallelNeighbors(11, 12))
		}

		if u.Cell(u.GetIndex(11, 13)) != Dead {
			t.Errorf("Expected cell %d to be dead, got %d", u.GetIndex(11, 12), u.Cell(u.GetIndex(11, 12)))
		}

		if u.ParallelNeighbors(11, 13) != 3 {
			t.Errorf("Expected cell %d to have 3 alive neighbors, got %d", u.GetIndex(11, 13), u.ParallelNeighbors(11, 13))
		}
	})
}

func TestNeighborData(t *testing.T) {
	t.Run("OneNeighbor", func(t *testing.T) {
		u := NewParallelUniverse(24, 32)
		u.cells[u.GetIndex(10, 12)] = Alive
		u.cells[u.GetIndex(11, 12)] = Alive
		u.cells[u.GetIndex(12, 12)] = Alive
		u.cells[u.GetIndex(11, 13)] = Alive

		u2 := NewParallelUniverse(24, 32)
		u2.cells[u.GetIndex(10, 12)] = Alive
		u2.cells[u.GetIndex(11, 12)] = Alive
		u2.cells[u.GetIndex(12, 12)] = Alive
		u2.cells[u.GetIndex(11, 13)] = Alive

		u.SetTopNeighbor(u2)

		u2.SendDataToNeighbors()
		u.WaitForNeighborsData()

		if len(u.TopNeighbor.Data.Data) != len(u2.cells) {
			t.Errorf("Expected data to be sent from top neighbor")
		}
	})

	t.Run("FourNeighbors", func(t *testing.T) {
		u := NewParallelUniverse(24, 32)
		u.cells[u.GetIndex(10, 12)] = Alive
		u.cells[u.GetIndex(11, 12)] = Alive
		u.cells[u.GetIndex(12, 12)] = Alive
		u.cells[u.GetIndex(11, 13)] = Alive

		u2 := NewParallelUniverse(24, 32)
		u2.cells[u.GetIndex(10, 12)] = Alive
		u2.cells[u.GetIndex(11, 12)] = Alive
		u2.cells[u.GetIndex(12, 12)] = Alive
		u2.cells[u.GetIndex(11, 13)] = Alive

		u3 := NewParallelUniverse(24, 32)
		u3.cells[u.GetIndex(10, 12)] = Alive
		u3.cells[u.GetIndex(11, 12)] = Alive
		u3.cells[u.GetIndex(12, 12)] = Alive
		u3.cells[u.GetIndex(11, 13)] = Alive

		u4 := NewParallelUniverse(24, 32)
		u4.cells[u.GetIndex(10, 12)] = Alive
		u4.cells[u.GetIndex(11, 12)] = Alive
		u4.cells[u.GetIndex(12, 12)] = Alive
		u4.cells[u.GetIndex(11, 13)] = Alive

		u5 := NewParallelUniverse(24, 32)
		u5.cells[u.GetIndex(10, 12)] = Alive
		u5.cells[u.GetIndex(11, 12)] = Alive
		u5.cells[u.GetIndex(12, 12)] = Alive
		u5.cells[u.GetIndex(11, 13)] = Alive

		u.SetTopNeighbor(u2)
		u.SetRightNeighbor(u3)
		u.SetBottomNeighbor(u4)
		u.SetLeftNeighbor(u5)

		go u2.SendDataToNeighbors()
		go u3.SendDataToNeighbors()
		go u4.SendDataToNeighbors()
		go u5.SendDataToNeighbors()
		u.WaitForNeighborsData()

		if len(u.TopNeighbor.Data.Data) != len(u2.cells) &&
			u.TopNeighbor.Data.Data[u.GetIndex(10, 12)] != Alive {
			t.Errorf("Expected data to be sent from top neighbor")
		}
		if len(u.RightNeighbor.Data.Data) != len(u3.cells) &&
			u.RightNeighbor.Data.Data[u.GetIndex(10, 12)] != Alive {
			t.Errorf("Expected data to be sent from right neighbor")
		}
		if len(u.BottomNeighbor.Data.Data) != len(u4.cells) &&
			u.BottomNeighbor.Data.Data[u.GetIndex(10, 12)] != Alive {
			t.Errorf("Expected data to be sent from bottom neighbor")
		}
		if len(u.LeftNeighbor.Data.Data) != len(u5.cells) &&
			u.LeftNeighbor.Data.Data[u.GetIndex(10, 12)] != Alive {
			t.Errorf("Expected data to be sent from left neighbor")
		}
	})
}

func TestMultitick(t *testing.T) {
	t.Run("FourNeighbors", func(t *testing.T) {
		u := NewParallelUniverse(24, 32)
		u.cells[u.GetIndex(10, 12)] = Alive
		u.cells[u.GetIndex(11, 12)] = Alive
		u.cells[u.GetIndex(12, 12)] = Alive
		u.cells[u.GetIndex(11, 13)] = Alive

		u2 := NewParallelUniverse(24, 32)
		u2.cells[u.GetIndex(10, 12)] = Alive
		u2.cells[u.GetIndex(11, 12)] = Alive
		u2.cells[u.GetIndex(12, 12)] = Alive
		u2.cells[u.GetIndex(11, 13)] = Alive

		u3 := NewParallelUniverse(24, 32)
		u3.cells[u.GetIndex(10, 12)] = Alive
		u3.cells[u.GetIndex(11, 12)] = Alive
		u3.cells[u.GetIndex(12, 12)] = Alive
		u3.cells[u.GetIndex(11, 13)] = Alive

		u4 := NewParallelUniverse(24, 32)
		u4.cells[u.GetIndex(10, 12)] = Alive
		u4.cells[u.GetIndex(11, 12)] = Alive
		u4.cells[u.GetIndex(12, 12)] = Alive
		u4.cells[u.GetIndex(11, 13)] = Alive

		u5 := NewParallelUniverse(24, 32)
		u5.cells[u.GetIndex(10, 12)] = Alive
		u5.cells[u.GetIndex(11, 12)] = Alive
		u5.cells[u.GetIndex(12, 12)] = Alive
		u5.cells[u.GetIndex(11, 13)] = Alive

		u.SetTopNeighbor(u2)
		u.SetRightNeighbor(u3)
		u.SetBottomNeighbor(u4)
		u.SetLeftNeighbor(u5)

		go u2.MultiTick()
		go u3.MultiTick()
		go u4.MultiTick()
		go u5.MultiTick()
		u.MultiTick()

		if u.TopNeighbor.Data.Data[u.GetIndex(10, 12)] != u2.cells[u.GetIndex(10, 12)] {
			t.Errorf("Expected data to be sent from top neighbor")
		}
		if u.RightNeighbor.Data.Data[u.GetIndex(10, 12)] != u3.cells[u.GetIndex(10, 12)] {
			t.Errorf("Expected data to be sent from right neighbor")
		}
		if u.BottomNeighbor.Data.Data[u.GetIndex(10, 12)] != u3.cells[u.GetIndex(10, 12)] {
			t.Errorf("Expected data to be sent from bottom neighbor")
		}
		if u.LeftNeighbor.Data.Data[u.GetIndex(10, 12)] != u4.cells[u.GetIndex(10, 12)] {
			t.Errorf("Expected data to be sent from left neighbor")
		}

		go u2.MultiTick()
		go u3.MultiTick()
		go u4.MultiTick()
		go u5.MultiTick()
		u.MultiTick()

		testCell(t, u)
		testCell(t, u2)
		testCell(t, u3)
		testCell(t, u4)
		testCell(t, u5)
	})
}

func testCell(t *testing.T, u *ParallelUniverse) {
	if u.Stable() {
		t.Errorf("Expected universe to not be stable")
	}

	if u.Cell(u.GetIndex(11, 12)) != Dead {
		t.Errorf("Expected cell %d to be dead, got %d", u.GetIndex(11, 12), u.Cell(u.GetIndex(11, 12)))
	}

	if u.ParallelNeighbors(11, 12) != 5 {
		t.Errorf("Expected cell %d to have 5 alive neighbors, got %d", u.GetIndex(11, 12), u.ParallelNeighbors(11, 12))
	}

	if u.Cell(u.GetIndex(11, 13)) != Dead {
		t.Errorf("Expected cell %d to be dead, got %d", u.GetIndex(11, 12), u.Cell(u.GetIndex(11, 12)))
	}

	if u.ParallelNeighbors(11, 13) != 3 {
		t.Errorf("Expected cell %d to have 3 alive neighbors, got %d", u.GetIndex(11, 13), u.ParallelNeighbors(11, 13))
	}
}
