package game

import (
	"crypto/rand"
)

func randomNumber() int {
	var b [1]byte
	rand.Read(b[:])
	val := int(b[0]) >> 1
	if val > 100 {
		val = val >> 1
	}
	return val
}
