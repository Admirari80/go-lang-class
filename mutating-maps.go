package main

// insert or update the element in map m
// m[key] = elem

// retrieve element
// elem = m[key]

// delete an element
// delete(m, key)

// test that a key is present with a two-value assignment
// elem, ok = m[key]
// if key is in m, ok is true. if not, ok is false

import "fmt"

func main() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The Value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The Value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The Value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The Value:", v, "Present?", ok)

}
