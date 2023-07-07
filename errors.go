package main

import (
	"errors"
	"fmt"
	"time"
)

// type error interface{
// 	Error()string
// }

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at indian time: %v, error msg: %s", e.When, e.What)
}

func run() error {
	return &MyError{time.Now(), "it didn't work"}
}

func Add(x, y int) (int, error) {
	if x > y {
		return x + y, nil
	}
	return x + y, errors.New("error in add func")
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
	sum, err := Add(20, 10)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		v, err := fmt.Println(sum)
		fmt.Println(v, err)
	}
}
