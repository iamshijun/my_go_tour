package main

import "fmt"

func main() {
	str := "\u4e2d\u5348\u597d"
	fmt.Println(len(str))
	fmt.Println(len([]rune(str)))

	var s string = "\u4e0b\u5348\u597d"
	fmt.Printf("Raw string : %s\n", s)
	fmt.Printf("Print with Range and %%U : ")
	for index, char := range s {
		fmt.Printf("% d:%U", index, char)
	}
	fmt.Println()

	fmt.Printf("Print with rune ,%%v and %%c : ")
	var rune_of_s = []rune(s)
	for i, len := 0, len(rune_of_s); i < len; i++ {
		fmt.Printf("% d:%v:%c ", i, rune_of_s[i], rune_of_s[i])
	}
	fmt.Println()

	var bytes_of_s = []byte(s)
	fmt.Printf("Byte Length of s : %d\n", len(bytes_of_s))
	for i, len := 0, len(bytes_of_s); i < len; i++ {
		fmt.Printf("% d:%X", i, bytes_of_s[i])
	}
	fmt.Println()
}
