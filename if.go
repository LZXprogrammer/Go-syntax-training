package main

import (
	"fmt"
	"math"
)

// if 的正常写法
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

// if 的简短写法
func pow(x, y, lim float64) float64 {

	if v := math.Pow(x, y); v < lim {
		return v
	}

	return lim
}


func main() {
	a, b := sqrt(2), sqrt(3)
	fmt.Println(a, b)

	fmt.Println(pow(2, 2, 5))
	fmt.Println(pow(3, 3, 20))
	// 或者组合写法
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}