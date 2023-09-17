package dmxconn

import (
	"fmt"
)

// Verify given channel ID is within range
func checkChannel(id int) (err error) {
	lowerBound := 1
	upperBound := 512
	if (id > upperBound) || (id < lowerBound) {
		err = fmt.Errorf("invalid channel [%d] - should be between '%d' and '%d'", id, lowerBound, upperBound)
	}
	return
}

// Verify given channel value is within range
func checkValue(val int) (err error) {
	lowerBound := 0
	upperBound := 255
	if (val > upperBound) || (val < lowerBound) {
		err = fmt.Errorf("invalid value [%d] - should be between '%d' and '%d'", val, lowerBound, upperBound)
	}
	return
}
