package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 25.0
	z -= (z*z - x) / (2 * z)
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	res, err := Sqrt(-2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
