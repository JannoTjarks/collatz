package collatz

import (
	"errors"
)

var ErrInputZeroToCollatz = errors.New("The number must be unequal zero")
var ErrInvalidInput = errors.New("The given input is not valid")
