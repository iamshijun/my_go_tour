package main

import (
	"fmt"
	"math"
)

//牛顿迭代法自己写平方根函数sqrt

func sqrt(x float64) float64 {
	for times := 1; times < 10; times++ {
		z = (z + x/z) / 2
	}
	return z
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println("Using math.Sqrt : ", math.Sqrt(2))
}
