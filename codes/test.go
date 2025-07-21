package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	str := fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
	return str
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		ErrNegativeSqrt(x).Error()
		return 0, ErrNegativeSqrt(x)
	}
	z := x + 1
	for {
		z -= (z*z - x) / (2 * z)
		if (z*z-x)*(z*z-x) < 0.001 {
			return z, nil
		}
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
