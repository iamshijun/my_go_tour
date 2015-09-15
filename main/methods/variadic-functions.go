package main

import "fmt"

//Variadic functions can be called with any number of trailing arguments.

func sum(nums ...int) int {
	fmt.Printf("Type of nums : %T , ", nums)
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {

	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))

	slice_nums := []int{1, 2, 3, 4}
	total := sum(slice_nums...) //!!!-> special use "..."
	fmt.Print(total)
}
