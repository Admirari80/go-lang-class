package main

import "fmt"

func PrintHere(data interface{}) {
	fmt.Println(data)
}

func main() {
	PrintHere(42)
	PrintHere(32.3)
	PrintHere("Ashish")
	PrintHere(true)
}
