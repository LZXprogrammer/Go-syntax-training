package main

import (
	"fmt"
	// "math"
)

func Sqrt(x float64) float64 {
	z := 1.0
    for i := 1.0; i < 8; i ++ {
		z -= (z*z - x) / (2*z)
		fmt.Println(z)
	}
	return 1.414
}

func main() {
    fmt.Println(Sqrt(4))
}