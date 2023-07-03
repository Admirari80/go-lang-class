package main

import "fmt"

func main() {
	names := [4]string{"Oliver", "Ashish", "John", "Rahul"}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	a[0] = "YYY"
	fmt.Println(a, b)  // [YYY, XXX], [XXX, john]
	fmt.Println(names) // [YYY, XXX, John, rahul]

}
