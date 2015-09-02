package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	//%q just like %s but string will be enbrace with " quote
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
