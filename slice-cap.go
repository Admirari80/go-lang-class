package main

import "fmt"

func main() {
	s := []int{}
	printSlice(s)
	for i := 0; i < 17; i++ {
		s = append(s, i)
		printSlice(s)
	}

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
