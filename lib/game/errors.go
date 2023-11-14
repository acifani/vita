package game

import (
	"errors"
)

var (
	errInvalidLength    = errors.New("slice does not match universe size")
	errInvalidID        = errors.New("IDs cannot be the same")
	errInvalidCharacter = errors.New("cannot parse invalid character")
)
