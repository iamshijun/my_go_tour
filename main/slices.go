package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)

	for i, length := 0, len(s); i < length; i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}
}
