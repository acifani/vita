package game

import (
	"testing"
)

func TestConwayRule(t *testing.T) {
	t.Run("MooreNeighbors", func(t *testing.T) {
		u := NewUniverse(32, 32)

		idx := u.GetIndex(12, 12)
		u.cells[idx] = Alive
		if u.Cell(idx) != Alive {
			t.Errorf("Expected cell %d to be alive, got %d", idx, u.Cell(idx))
		}

		if u.MooreNeighbors(0, 0) != 0 {
			t.Errorf("Expected cell %d to have 0 alive neighbors, got %d", 0, u.MooreNeighbors(0, 0))
		}

		if u.MooreNeighbors(11, 12) != 1 {
			t.Errorf("Expected cell %d to have 1 alive neighbors, got %d", idx, u.MooreNeighbors(11, 12))
		}

		idx = u.GetIndex(10, 12)
		u.cells[idx] = Alive

		if u.MooreNeighbors(11, 12) != 2 {
			t.Errorf("Expected cell %d to have 2 alive neighbors, got %d", idx, u.MooreNeighbors(11, 12))
		}
	})

	t.Run("MooreNeighborsWrap", func(t *testing.T) {
		u := NewUniverse(32, 32)

		idx := u.GetIndex(12, 12)
		u.cells[idx] = Alive
		if u.Cell(idx) != Alive {
			t.Errorf("Expected cell %d to be alive, got %d", idx, u.Cell(idx))
		}

		if u.MooreNeighborsWrap(0, 0) != 0 {
			t.Errorf("Expected cell %d to have 0 alive neighbors, got %d", 0, u.MooreNeighborsWrap(0, 0))
		}

		if u.MooreNeighborsWrap(11, 12) != 1 {
			t.Errorf("Expected cell %d to have 1 alive neighbors, got %d", idx, u.MooreNeighborsWrap(11, 12))
		}

		idx = u.GetIndex(10, 12)
		u.cells[idx] = Alive

		if u.MooreNeighborsWrap(11, 12) != 2 {
			t.Errorf("Expected cell %d to have 2 alive neighbors, got %d", idx, u.MooreNeighborsWrap(11, 12))
		}
	})

	t.Run("RuleB3S23 when cell is Alive", func(t *testing.T) {
		if RuleB3S23(Alive, 0) != Dead {
			t.Errorf("Expected cell to be dead, got %d", RuleB3S23(Alive, 0))
		}

		if RuleB3S23(Alive, 1) != Dead {
			t.Errorf("Expected cell to be dead, got %d", RuleB3S23(Alive, 1))
		}

		if RuleB3S23(Alive, 2) != Alive {
			t.Errorf("Expected cell to be alive, got %d", RuleB3S23(Alive, 2))
		}

		if RuleB3S23(Alive, 3) != Alive {
			t.Errorf("Expected cell to be alive, got %d", RuleB3S23(Alive, 3))
		}

		if RuleB3S23(Alive, 4) != Dead {
			t.Errorf("Expected cell to be dead, got %d", RuleB3S23(Alive, 4))
		}
	})

	t.Run("RuleB3S23 when cell is Dead", func(t *testing.T) {
		if RuleB3S23(Dead, 0) != Dead {
			t.Errorf("Expected cell to be dead, got %d", RuleB3S23(Dead, 0))
		}

		if RuleB3S23(Dead, 1) != Dead {
			t.Errorf("Expected cell to be dead, got %d", RuleB3S23(Dead, 1))
		}

		if RuleB3S23(Dead, 2) != Dead {
			t.Errorf("Expected cell to be dead, got %d", RuleB3S23(Dead, 2))
		}

		if RuleB3S23(Dead, 3) != Alive {
			t.Errorf("Expected cell to be alive, got %d", RuleB3S23(Dead, 3))
		}

		if RuleB3S23(Dead, 4) != Dead {
			t.Errorf("Expected cell to be dead, got %d", RuleB3S23(Dead, 4))
		}
	})

	t.Run("RuleB3678S34678 when cell is Alive", func(t *testing.T) {
		if RuleB3678S34678(Alive, 0) != Dead {
			t.Errorf("Expected cell to be dead, got %d", RuleB3678S34678(Alive, 0))
		}

		if RuleB3678S34678(Alive, 1) != Dead {
			t.Errorf("Expected cell to be dead, got %d", RuleB3678S34678(Alive, 1))
		}

		if RuleB3678S34678(Alive, 2) != Dead {
			t.Errorf("Expected cell to be dead, got %d", RuleB3678S34678(Dead, 2))
		}

		if RuleB3678S34678(Alive, 3) != Alive {
			t.Errorf("Expected cell to be alive, got %d", RuleB3678S34678(Alive, 3))
		}

		if RuleB3678S34678(Alive, 4) != Alive {
			t.Errorf("Expected cell to be alive, got %d", RuleB3678S34678(Alive, 4))
		}

		if RuleB3678S34678(Alive, 5) != Dead {
			t.Errorf("Expected cell to be dead, got %d", RuleB3678S34678(Alive, 5))
		}

		if RuleB3678S34678(Alive, 6) != Alive {
			t.Errorf("Expected cell to be alive, got %d", RuleB3678S34678(Alive, 6))
		}

		if RuleB3678S34678(Alive, 7) != Alive {
			t.Errorf("Expected cell to be alive, got %d", RuleB3678S34678(Alive, 7))
		}

		if RuleB3678S34678(Alive, 8) != Alive {
			t.Errorf("Expected cell to be alive, got %d", RuleB3678S34678(Alive, 8))
		}

		if RuleB3678S34678(Alive, 9) != Dead {
			t.Errorf("Expected cell to be dead, got %d", RuleB3678S34678(Alive, 9))
		}
	})

	t.Run("RuleB3678S34678 when cell is Dead", func(t *testing.T) {
		if RuleB3678S34678(Dead, 0) != Dead {
			t.Errorf("Expected cell to be dead, got %d", RuleB3678S34678(Dead, 0))
		}

		if RuleB3678S34678(Dead, 1) != Dead {
			t.Errorf("Expected cell to be dead, got %d", RuleB3678S34678(Dead, 1))
		}

		if RuleB3678S34678(Dead, 2) != Dead {
			t.Errorf("Expected cell to be dead, got %d", RuleB3678S34678(Dead, 2))
		}

		if RuleB3678S34678(Dead, 3) != Alive {
			t.Errorf("Expected cell to be alive, got %d", RuleB3678S34678(Dead, 3))
		}

		if RuleB3678S34678(Dead, 4) != Dead {
			t.Errorf("Expected cell to be alive, got %d", RuleB3678S34678(Dead, 4))
		}

		if RuleB3678S34678(Dead, 5) != Dead {
			t.Errorf("Expected cell to be dead, got %d", RuleB3678S34678(Dead, 5))
		}

		if RuleB3678S34678(Dead, 6) != Alive {
			t.Errorf("Expected cell to be alive, got %d", RuleB3678S34678(Dead, 6))
		}

		if RuleB3678S34678(Dead, 7) != Alive {
			t.Errorf("Expected cell to be alive, got %d", RuleB3678S34678(Dead, 7))
		}

		if RuleB3678S34678(Dead, 8) != Alive {
			t.Errorf("Expected cell to be alive, got %d", RuleB3678S34678(Dead, 8))
		}

		if RuleB3678S34678(Dead, 9) != Dead {
			t.Errorf("Expected cell to be dead, got %d", RuleB3678S34678(Dead, 9))
		}
	})
}
