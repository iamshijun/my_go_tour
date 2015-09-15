package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func main() {
	a := make([]int, 5) //[0 0 0 0 0]
	fmt.Printf("type of slice : %T\n", a)
	printSlice("a", a)

	b := make([]int, 0, 5) // len=0 , cap=5 [] 0 0 0 0 0
	t_b := b[:cap(b)]
	for i, len_of_t_b := 0, len(t_b); i < len_of_t_b; i++ {
		t_b[i] = i
	}
	printSlice("b", b)
	//fmt.Println(b[2])

	c := b[:2] //len=2 cap=5 [0 1] 2 3 4
	printSlice("c", c)

	d := c[2:5] //len=3 cap=3 [2 3 4]
	printSlice("d", d)

	//additional test
	b = make([]int, 0, 5) // len=0 , cap=5 [] 0 0 0 0 0
	b = b[:cap(b)]        // b[0:5] , len(b)=5,cap(5) [0 0 0 0 0]
	printSlice("redefine b", b)

	b = b[1:] // len(b)=4,cap(4)
	printSlice("redefine b", b)
	fmt.Println(b[2])
}
