package main

import "fmt"

func fibonacci() func() int {
	v1, v2 := 0, 1
	return func() int {
		v1, v2 = v2, v1+v2
		return v1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

//yosi
