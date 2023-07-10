package main

// v , ok := <-ch
// close(ch)

import "fmt"

func fibonacci(n int64, c chan int64) {
	var x, y int64 = 0, 1
	var i int64
	for i = 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int64, 10)
	go fibonacci(100, c)
	for i := range c {
		fmt.Println(i)
	}
}
