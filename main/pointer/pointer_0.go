package main

import "fmt"

func main() {
	var pp **int
	fmt.Println(pp)
	i := 1
	pi := &i
	ppi := &pi

	//ppi = &(&i) //wrong
	fmt.Println(ppi)

	var arr = [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%T \n", arr)
	pa := &arr
	fmt.Printf("type of p : %T , pointer(addr) : %p , val : %v\n", pa, pa, pa)
	fmt.Printf("%-10s %v\n", "*pa", *pa)
	fmt.Printf("%-10s %v\n", "pa[2]", pa[2])
	fmt.Printf("%-10s %v\n", "(*pa)[2]", (*pa)[2])

	//pointer to array , you can use it just like array
	//C program  get arr[1] <= *(*pa+1) or (*pa)[1]
	//Go  pa[1]

}
