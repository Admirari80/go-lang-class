package main

import "fmt"

// A pointer holds the memory address of a value

// *T is a pointer to a T value zero value nil

// var p *int

// i := 42
// p = &i

func main() {
	i, j := 42, 2701

	k := i
	fmt.Println(k)
	k = k + 10
	fmt.Println(k)
	fmt.Println(i)

	p := &i
	fmt.Println(p)  // memory address
	fmt.Println(*p) // 42
	*p = 21
	fmt.Println(i) // 21

	p = &j
	*p = *p / 37
	fmt.Println(p)
	fmt.Println(j) // 73
}
