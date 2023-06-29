package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	t = t.Add(18 * time.Hour)
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
