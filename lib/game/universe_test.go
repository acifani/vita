package game

import (
	"testing"
)

func TestUniverse(t *testing.T) {
	t.Run("NewUniverse", func(t *testing.T) {
		u := NewUniverse(24, 32)

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

	t.Run("AliveNeighbors", func(t *testing.T) {
		u := NewUniverse(32, 32)

		idx := u.GetIndex(12, 12)
		u.cells[idx] = Alive
		if u.Cell(idx) != Alive {
			t.Errorf("Expected cell %d to be alive, got %d", idx, u.Cell(idx))
		}

		if u.AliveNeighbors(0, 0) != 0 {
			t.Errorf("Expected cell %d to have 0 alive neighbors, got %d", 0, u.AliveNeighbors(0, 0))
		}

		if u.AliveNeighbors(11, 12) != 1 {
			t.Errorf("Expected cell %d to have 1 alive neighbors, got %d", idx, u.AliveNeighbors(11, 12))
		}

		idx = u.GetIndex(10, 12)
		u.cells[idx] = Alive

		if u.AliveNeighbors(11, 12) != 2 {
			t.Errorf("Expected cell %d to have 2 alive neighbors, got %d", idx, u.AliveNeighbors(11, 12))
		}
	})

	t.Run("Tick", func(t *testing.T) {
		u := NewUniverse(24, 32)
		u.cells[u.GetIndex(10, 12)] = Alive
		u.cells[u.GetIndex(11, 12)] = Alive
		u.cells[u.GetIndex(12, 12)] = Alive
		u.cells[u.GetIndex(11, 13)] = Alive

		u.Tick()

		if u.AliveNeighbors(0, 0) != 0 {
			t.Errorf("Expected cell %d to have 0 alive neighbors, got %d", 0, u.AliveNeighbors(0, 0))
		}

		if u.AliveNeighbors(11, 12) != 6 {
			t.Errorf("Expected cell %d to have 6 alive neighbors, got %d", 0, u.AliveNeighbors(11, 12))
		}

		u.Tick()

		if u.Cell(u.GetIndex(11, 12)) != Dead {
			t.Errorf("Expected cell %d to be dead, got %d", u.GetIndex(11, 12), u.Cell(u.GetIndex(11, 12)))
		}

		if u.AliveNeighbors(11, 12) != 5 {
			t.Errorf("Expected cell %d to have 5 alive neighbors, got %d", u.GetIndex(11, 12), u.AliveNeighbors(11, 12))
		}

		if u.Cell(u.GetIndex(11, 13)) != Dead {
			t.Errorf("Expected cell %d to be dead, got %d", u.GetIndex(11, 12), u.Cell(u.GetIndex(11, 12)))
		}

		if u.AliveNeighbors(11, 13) != 3 {
			t.Errorf("Expected cell %d to have 3 alive neighbors, got %d", u.GetIndex(11, 13), u.AliveNeighbors(11, 13))
		}
	})

	t.Run("Rule when cell is Alive", func(t *testing.T) {
		if rule(Alive, 0) != Dead {
			t.Errorf("Expected cell to be dead, got %d", rule(Alive, 0))
		}

		if rule(Alive, 1) != Dead {
			t.Errorf("Expected cell to be dead, got %d", rule(Alive, 1))
		}

		if rule(Alive, 2) != Alive {
			t.Errorf("Expected cell to be alive, got %d", rule(Alive, 2))
		}

		if rule(Alive, 3) != Alive {
			t.Errorf("Expected cell to be alive, got %d", rule(Alive, 3))
		}

		if rule(Alive, 4) != Dead {
			t.Errorf("Expected cell to be dead, got %d", rule(Alive, 4))
		}
	})

	t.Run("Rule when cell is Dead", func(t *testing.T) {
		if rule(Dead, 0) != Dead {
			t.Errorf("Expected cell to be dead, got %d", rule(Dead, 0))
		}

		if rule(Dead, 1) != Dead {
			t.Errorf("Expected cell to be dead, got %d", rule(Dead, 1))
		}

		if rule(Dead, 2) != Dead {
			t.Errorf("Expected cell to be dead, got %d", rule(Dead, 2))
		}

		if rule(Dead, 3) != Alive {
			t.Errorf("Expected cell to be alive, got %d", rule(Dead, 3))
		}

		if rule(Dead, 4) != Dead {
			t.Errorf("Expected cell to be dead, got %d", rule(Dead, 4))
		}
	})
}
