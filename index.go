package main

import "fmt"

func Index[T any](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i + 1
		}
	}
	return -1
}

func main() {
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	ss := []string{"foo", "bar", "baz", "Ashish", "oliver"}
	fmt.Println(Index(ss, "oliver"))
}
