package main

import (
	"fmt"
	"math"
)

type Vehicle interface {
	Milage() float64
}

type Car struct {
	distance float64
	fuel     float64
}

func (c Car) Milage() float64 {
	return c.distance / c.fuel
}

type Truck struct {
	distance     float64
	fuel         float64
	AverageSpeed float64
}

func (t Truck) Milage() float64 {
	return t.distance / t.fuel
}

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	fmt.Println(a.Abs())

	var vi Vehicle
	car := Car{125, 12}
	truck := Truck{100, 20, 40}

	vi = car
	fmt.Println(vi.Milage())
	vi = truck
	fmt.Println(vi.Milage())

}
