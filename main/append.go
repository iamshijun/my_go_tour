package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%-5s %p len=%d cap=%d %v\n", s, &x, len(x), cap(x), x)
}

/**
http://blog.go-zh.org/go-slices-usage-and-internals
*/
func main() {
	arr := [...]int{1, 3, 5, 7, 9}
	fmt.Printf("%-5s %p\n", "arr", &arr)

	s_arr := arr[:cap(arr)]
	fmt.Printf("%-5s %p\n", "s_arr", &s_arr)
	//printSlice("arr", arr)
	/*
		cannot use arr (type [5]int) as type []int in argument to printSlice
		the sencond param of printSlice is "Slice" not array,if you want to pass arr ,
		you have to change the param's type into [5]int explicitly tell method want arr array with 5 element
	*/

	var a []int //slice a with default zero value => nil
	//fmt.Printf("%p\n", &a)
	printSlice("a", a)

	//build-in func append
	//func append(s []T, vs ...T) []T
	a = append(a, 0)
	printSlice("a", a)

	a = append(a, 1)
	printSlice("a", a)

	a = append(a, 2, 3, 4)
	printSlice("a", a)

	a = append(a, 5, 6, 7) //grow twice of old cap
	printSlice("a", a)

}
