package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	//fmt.Printf("+%p+\n", v) //same as line30 variable "v"
}

/*
func (v Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Printf("+%p+\n", v)
}*/

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	//fmt.Printf("-%p-\n", v)
	v.Scale(5)
	fmt.Println(v, v.Abs())
}
