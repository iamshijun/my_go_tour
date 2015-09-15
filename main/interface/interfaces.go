package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

//!类型通过实现那些方法来实现接口。 没有显式声明的必要；所以也就没有关键字“implements“。

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	fmt.Println(a.Abs())
	//a = v //here we will go error
	//Vertex does not implement Abser (Abs method has pointer receiver)
	//cause only *Vertex implement Abser not Vertex
	a = &v
	fmt.Println(a.Abs())
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
