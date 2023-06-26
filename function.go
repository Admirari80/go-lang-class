package main

import(
	"fmt"
)

func add(x,y int)int{
	return x+y
}

func swap(x, y string) (string, string){
    return y ,x
}

func operation(x,y int)(int, int){
	return x+y, x*y
}

func main(){
	fmt.Println(add(23,56))
	a, b := swap("hello", "world")
	fmt.Println(a,b)
	fmt.Println(operation(10, 20))
}
