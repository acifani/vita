package game

import (
	"testing"
)

func TestRandomness(t *testing.T) {
	val1 := randomNumber()
	val2 := randomNumber()
	val3 := randomNumber()

	if val1 == val2 || val1 == val3 || val2 == val3 {
		t.Errorf("Expected %d and %d and %d to be different", val1, val2, val3)
	}
}
