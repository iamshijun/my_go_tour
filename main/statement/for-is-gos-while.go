package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 { // while(sum < 1000)
		//go does not have "while" keyword, cause for can do what it can do
		sum += sum
	}
	fmt.Println(sum)

	//forever
	/*
		for {
		}

		   just like

		while(true) {
		}
	*/
}
