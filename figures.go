package main

import (
	"github.com/acifani/vita/lib/game"
)

type figure struct {
	deltaX uint32
	deltaY uint32
	values [][]uint8
}

// Inspired by: https://www.reddit.com/r/rust/comments/5penft/comment/dcsq64p
// If you look closely, those aren't angle brackets,
// they're characters from the Canadian Aboriginal Syllabics block,
// which are allowed in Go identifiers.
const (
	ᑕᑐ = game.Dead
	ᕳᕲ = game.Alive
)

func glider() *figure {
	return &figure{
		deltaX: 1,
		deltaY: 1,
		values: [][]uint8{
			{ᑕᑐ, ᕳᕲ, ᑕᑐ},
			{ᑕᑐ, ᑕᑐ, ᕳᕲ},
			{ᕳᕲ, ᕳᕲ, ᕳᕲ},
		},
	}
}

func pulsar() *figure {
	return &figure{
		deltaX: 6,
		deltaY: 6,
		values: [][]uint8{
			{ᑕᑐ, ᑕᑐ, ᕳᕲ, ᕳᕲ, ᕳᕲ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᕳᕲ, ᕳᕲ, ᕳᕲ, ᑕᑐ, ᑕᑐ},
			{ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ},
			{ᕳᕲ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᕳᕲ, ᑕᑐ, ᕳᕲ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᕳᕲ},
			{ᕳᕲ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᕳᕲ, ᑕᑐ, ᕳᕲ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᕳᕲ},
			{ᕳᕲ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᕳᕲ, ᑕᑐ, ᕳᕲ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᕳᕲ},
			{ᑕᑐ, ᑕᑐ, ᕳᕲ, ᕳᕲ, ᕳᕲ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᕳᕲ, ᕳᕲ, ᕳᕲ, ᑕᑐ, ᑕᑐ},
			{ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ},
			{ᑕᑐ, ᑕᑐ, ᕳᕲ, ᕳᕲ, ᕳᕲ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᕳᕲ, ᕳᕲ, ᕳᕲ, ᑕᑐ, ᑕᑐ},
			{ᕳᕲ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᕳᕲ, ᑕᑐ, ᕳᕲ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᕳᕲ},
			{ᕳᕲ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᕳᕲ, ᑕᑐ, ᕳᕲ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᕳᕲ},
			{ᕳᕲ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᕳᕲ, ᑕᑐ, ᕳᕲ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᕳᕲ},
			{ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᑕᑐ},
			{ᑕᑐ, ᑕᑐ, ᕳᕲ, ᕳᕲ, ᕳᕲ, ᑕᑐ, ᑕᑐ, ᑕᑐ, ᕳᕲ, ᕳᕲ, ᕳᕲ, ᑕᑐ, ᑕᑐ},
		},
	}
}
