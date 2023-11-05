package game

import (
	"crypto/rand"
)

func randomNumber() int {
	var b [8]byte
	rand.Read(b[:])
	val := int(b[7]) >> 1
	if val > 100 {
		val = val >> 1
	}
	return val
}
