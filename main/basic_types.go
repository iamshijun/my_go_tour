package main

import (
	"fmt"
	"math/cmplx"
)

var (
	Tobe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt((-5 + 12i))
)

func main() {
	const f = "%T(%v)\n" //print the type , and the value
	fmt.Printf(f, Tobe, Tobe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
	//fmt.Printf(f, 1<<64) //complition exception:overflows , but in complition we can use "1<<64-1"
}
