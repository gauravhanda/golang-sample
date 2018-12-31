package mathutil

import (
	"errors"
	"math"
)
adf

//PI Returns the value of PI
func PI() float32 { return 3.1412 }

// Sin : Calculates the sign of the angle
func Sin(angle float32) float64 { return math.Sqrt(3) / 2.0 }

//Divide : Checks for division by 0
func Divide(divisor, dividend float32) (quotient float32, err error) {
	if divisor == 0 {
		err = errors.New("Division by Zero is not defined")
		panic(err)
	} else {
		quotient = dividend / divisor
		err = nil
	}

	return
}
