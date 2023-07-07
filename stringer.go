package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
func main() {
	a := Person{"Oliver Wunder", 42}
	z := Person{"Ashish Singh", 25}
	fmt.Println(a, z)
}
