package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Nextafter(2, 3))
	fmt.Println()

	fmt.Println("math.Pi is : ", math.Pi)
}
