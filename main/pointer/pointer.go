package main

import "fmt"

func main() {
	i, j := 42, 7201

	pi := &i
	var pj *int = &j

	fmt.Println(*pi)
	*pi = 21
	fmt.Println("i == 21 :", i == 21)

	*pj = *pj / 37
	fmt.Println("*pj = *pj / 37 : ", j)
	//pointer in go look like the same as C but poninter in go cannot do calculate
}
