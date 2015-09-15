package main

import "fmt"

func main() {
	//var a []int
	a := [...]int{1, 2, 3, 4, 5} //let complier set the num of a => [5]int{1,2,3,4,5}
	sa := a
	fmt.Println(sa)

	pa := &a //type of pa => *[5]int (pointer to []int)
	//so  *pa is an array => [5]int
	fmt.Println(*pa)   // value of "*pa" is equal to "a"
	fmt.Println(pa[0]) // 1
}
