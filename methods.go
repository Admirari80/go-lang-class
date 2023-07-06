package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) ScalerMethod(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScalerFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.ScalerMethod(10)
	(&v).ScalerMethod(10)
	ScalerFunc(&v, 10)

	p := &Vertex{4, 3}
	p.ScalerMethod(3)
	ScalerFunc(p, 8)

	fmt.Println(v, p)

}
