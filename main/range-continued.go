package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	fmt.Println(pow)

	pow_2_5 := pow[2:5]
	for i, _ := range pow_2_5 {
		fmt.Printf("% d", i)
	}
	fmt.Println()

	pow_5_ := pow[5:]
	for _, v := range pow_5_ {
		fmt.Printf("% d", v)
	}

}
