package main

import "fmt"

func generic_circumference[r int | float32](radius r) {
	c := 2 * 3 * radius
	fmt.Println("the circumferece is: ", c)
}

func Swap[s int | float64 | string](x, y s) (s, s) {
	return y, x
}

// func SwapInt(x, y int) (int, int) {
// 	return y, x
// }

// func SwapFloat(x, y float64) (float64, float64) {
// 	return y, x
// }

// func SwapString(x, y string) (string, string) {
// 	return y, x
// }

func main() {
	var rad1 int = 8
	var rad2 float32 = 9.5

	generic_circumference(rad1)
	generic_circumference(rad2)

	fmt.Println(Swap(10, 20))
	fmt.Println(Swap(10.1, 20.2))
	fmt.Println(Swap("Ashish", "oliver"))
}
