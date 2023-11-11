package game

import (
	"testing"
)

func TestDistributedUniverse(t *testing.T) {
	t.Run("NewDistributedUniverse", func(t *testing.T) {
		u := NewDistributedUniverse(GenerateKey(), 24, 32)

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

		if len(u.Bytes()) != 24*32+5*IDLength {
			t.Errorf("Expected Bytes to be %d, got %d", 24*32+5*IDLength, len(u.Bytes()))
		}
	})

	t.Run("TickWithoutNeighbors", func(t *testing.T) {
		u := NewDistributedUniverse(GenerateKey(), 24, 32)
		u.cells[u.GetIndex(10, 12)] = Alive
		u.cells[u.GetIndex(11, 12)] = Alive
		u.cells[u.GetIndex(12, 12)] = Alive
		u.cells[u.GetIndex(11, 13)] = Alive

		u.Tick()

		if u.Stable() {
			t.Errorf("Expected universe to not be stable")
		}

		if u.Neighbors(0, 0) != 0 {
			t.Errorf("Expected cell %d to have 0 alive neighbors, got %d", u.GetIndex(0, 0), u.Neighbors(0, 0))
		}

		if u.Neighbors(11, 12) != 6 {
			t.Errorf("Expected cell %d to have 6 alive neighbors, got %d", u.GetIndex(11, 12), u.Neighbors(11, 12))
		}

		u.Tick()

		if u.Stable() {
			t.Errorf("Expected universe to not be stable")
		}

		if u.Cell(u.GetIndex(11, 12)) != Dead {
			t.Errorf("Expected cell %d to be dead, got %d", u.GetIndex(11, 12), u.Cell(u.GetIndex(11, 12)))
		}

		if u.Neighbors(11, 12) != 5 {
			t.Errorf("Expected cell %d to have 5 alive neighbors, got %d", u.GetIndex(11, 12), u.Neighbors(11, 12))
		}

		if u.Cell(u.GetIndex(11, 13)) != Dead {
			t.Errorf("Expected cell %d to be dead, got %d", u.GetIndex(11, 12), u.Cell(u.GetIndex(11, 12)))
		}

		if u.Neighbors(11, 13) != 3 {
			t.Errorf("Expected cell %d to have 3 alive neighbors, got %d", u.GetIndex(11, 13), u.Neighbors(11, 13))
		}
	})
}

func TestDistributedUniverseTick(t *testing.T) {
	t.Run("FourNeighbors", func(t *testing.T) {
		store := NewTestStore()

		u := NewDistributedUniverse(GenerateKey(), 24, 32)
		u.GetNeighbor = store.Get
		store.Set(u.ID, u)
		u.cells[u.GetIndex(10, 12)] = Alive
		u.cells[u.GetIndex(11, 12)] = Alive
		u.cells[u.GetIndex(12, 12)] = Alive
		u.cells[u.GetIndex(11, 13)] = Alive

		u2 := NewDistributedUniverse(GenerateKey(), 24, 32)
		u2.GetNeighbor = store.Get
		store.Set(u2.ID, u2)
		u2.cells[u.GetIndex(10, 12)] = Alive
		u2.cells[u.GetIndex(11, 12)] = Alive
		u2.cells[u.GetIndex(12, 12)] = Alive
		u2.cells[u.GetIndex(11, 13)] = Alive

		u3 := NewDistributedUniverse(GenerateKey(), 24, 32)
		u3.GetNeighbor = store.Get
		store.Set(u3.ID, u3)
		u3.cells[u.GetIndex(10, 12)] = Alive
		u3.cells[u.GetIndex(11, 12)] = Alive
		u3.cells[u.GetIndex(12, 12)] = Alive
		u3.cells[u.GetIndex(11, 13)] = Alive

		u4 := NewDistributedUniverse(GenerateKey(), 24, 32)
		u4.GetNeighbor = store.Get
		store.Set(u4.ID, u4)
		u4.cells[u.GetIndex(10, 12)] = Alive
		u4.cells[u.GetIndex(11, 12)] = Alive
		u4.cells[u.GetIndex(12, 12)] = Alive
		u4.cells[u.GetIndex(11, 13)] = Alive

		u5 := NewDistributedUniverse(GenerateKey(), 24, 32)
		u5.GetNeighbor = store.Get
		store.Set(u5.ID, u5)
		u5.cells[u.GetIndex(10, 12)] = Alive
		u5.cells[u.GetIndex(11, 12)] = Alive
		u5.cells[u.GetIndex(12, 12)] = Alive
		u5.cells[u.GetIndex(11, 13)] = Alive

		u.SetTopNeighbor(u2)
		u2.SetBottomNeighbor(u)
		u.SetRightNeighbor(u3)
		u3.SetLeftNeighbor(u)
		u.SetBottomNeighbor(u4)
		u4.SetTopNeighbor(u)
		u.SetLeftNeighbor(u5)
		u5.SetRightNeighbor(u)

		u.Tick()
		u2.Tick()
		u3.Tick()
		u4.Tick()
		u5.Tick()

		if u.TopNeighbor.Cell(u.GetIndex(10, 12)) != u2.Cell(u.GetIndex(10, 12)) {
			t.Errorf("Expected data to be sent from top neighbor")
		}
		if u.RightNeighbor.Cell(u.GetIndex(10, 12)) != u3.Cell(u.GetIndex(10, 12)) {
			t.Errorf("Expected data to be sent from right neighbor")
		}
		if u.BottomNeighbor.Cell(u.GetIndex(10, 12)) != u3.Cell(u.GetIndex(10, 12)) {
			t.Errorf("Expected data to be sent from bottom neighbor")
		}
		if u.LeftNeighbor.Cell(u.GetIndex(10, 12)) != u4.Cell(u.GetIndex(10, 12)) {
			t.Errorf("Expected data to be sent from left neighbor")
		}

		u.Tick()
		u2.Tick()
		u3.Tick()
		u4.Tick()
		u5.Tick()

		testDistributedCell(t, u)
		testDistributedCell(t, u2)
		testDistributedCell(t, u3)
		testDistributedCell(t, u4)
		testDistributedCell(t, u5)
	})
}

func testDistributedCell(t *testing.T, u *DistributedUniverse) {
	if u.Stable() {
		t.Errorf("Expected universe to not be stable")
	}

	if u.Cell(u.GetIndex(11, 12)) != Dead {
		t.Errorf("Expected cell %d to be dead, got %d", u.GetIndex(11, 12), u.Cell(u.GetIndex(11, 12)))
	}

	if u.Neighbors(11, 12) != 5 {
		t.Errorf("Expected cell %d to have 5 alive neighbors, got %d", u.GetIndex(11, 12), u.Neighbors(11, 12))
	}

	if u.Cell(u.GetIndex(11, 13)) != Dead {
		t.Errorf("Expected cell %d to be dead, got %d", u.GetIndex(11, 12), u.Cell(u.GetIndex(11, 12)))
	}

	if u.Neighbors(11, 13) != 3 {
		t.Errorf("Expected cell %d to have 3 alive neighbors, got %d", u.GetIndex(11, 13), u.Neighbors(11, 13))
	}
}

type testStore struct {
	data map[string]*DistributedUniverse
}

func NewTestStore() *testStore {
	return &testStore{data: make(map[string]*DistributedUniverse)}
}

func (t *testStore) Set(key string, val *DistributedUniverse) {
	t.data[key] = val
}

func (t *testStore) Get(key string) *DistributedUniverse {
	return t.data[key]
}
