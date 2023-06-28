package main

import "fmt"

func main() {
	sum := 0
	var i int
	for i = 0; i < 21; i++ {
		sum += i
	}
	fmt.Println(sum)
}
