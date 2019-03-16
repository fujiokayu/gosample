package gosample

import (
	"errors"
)

func div(i, j int) (int, error) {
	if j == 0 {
		// return my error
		return 0, errors.New("divied by zero")
	}
	return i / j, nil
}
