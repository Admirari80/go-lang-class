package main

// go f(x,y,z)

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	//say("hello")
	time.Sleep(550 * time.Millisecond)
}
