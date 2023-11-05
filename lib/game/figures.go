package game

type Figure struct {
	deltaX uint32
	deltaY uint32
	values [][]uint8
}

func (f *Figure) DeltaX() uint32 {
	return f.deltaX
}

func (f *Figure) DeltaY() uint32 {
	return f.deltaY
}

func (f *Figure) Values() [][]uint8 {
	return f.values
}

// Inspired by: https://www.reddit.com/r/rust/comments/5penft/comment/dcsq64p
// If you look closely, those aren't angle brackets,
// they're characters from the Canadian Aboriginal Syllabics block,
// which are allowed in Go identifiers.
const (
	ᑕᑐ = Dead
	ᕳᕲ = Alive
)

func Glider() *Figure {
	return &Figure{
		deltaX: 1,
		deltaY: 1,
		values: [][]uint8{
			{ᑕᑐ, ᕳᕲ, ᑕᑐ},
			{ᑕᑐ, ᑕᑐ, ᕳᕲ},
			{ᕳᕲ, ᕳᕲ, ᕳᕲ},
		},
	}
}

func Pulsar() *Figure {
	return &Figure{
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

func Beehive() *Figure {
	return &Figure{
		deltaX: 1,
		deltaY: 1,
		values: [][]uint8{
			{ᑕᑐ, ᕳᕲ, ᕳᕲ, ᑕᑐ},
			{ᕳᕲ, ᑕᑐ, ᑕᑐ, ᕳᕲ},
			{ᑕᑐ, ᕳᕲ, ᕳᕲ, ᑕᑐ},
		},
	}
}
