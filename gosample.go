package gosample

import (
	"errors"
)

type firstVal int
type secondVal int

func Div(i firstVal, j secondVal) (int, error) {
	if j == 0 {
		// return my error
		return 0, errors.New("divied by zero")
	}

	return i / j, nil
}
