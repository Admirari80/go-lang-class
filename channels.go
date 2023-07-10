package main

import "fmt"

// <-
// ch <- v // send v to channel ch
// v := <-ch // receive from ch and assign value to v
// ch := make(chan int)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	ch := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], ch)
	x, y := <-c, <-ch // recive from c

	fmt.Println(x, y, x+y)
}
