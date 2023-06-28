package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func test(x int) int {
	if x < 10 {
		fmt.Println("below 10")
		return x
	} else if x > 10 && x < 20 {
		fmt.Println("above 10 and below 20")
		return x
	} else {
		fmt.Println("above 20")
		return x
	}
}

func main() {
	var i int = 100
	if i < 20 {
		fmt.Println("small number")
	}
	if i > 20 {
		fmt.Println("big number")
	}

	fmt.Println(pow(3, 2, 10), pow(3, 3, 10))

	fmt.Println(test(5), test(11), test(21))
}

// find the sum of the square root of 1 to 100
