package main

import "fmt"

func main() {
	s := "\u4e2d\u5348\u597d"
	fmt.Println(len(s))
	fmt.Println(len([]rune(s)))
}
