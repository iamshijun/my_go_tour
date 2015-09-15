package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	//for 循环的 range 格式可以对 slice / array数组 或者 map 进行迭代循环。
	//还有前面讲述到的 string 返回的v是每一个utf-8字符
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	a := [...]int{1, 3, 5, 7, 9}
	for i, v := range a {
		fmt.Printf("% d:%d", i, v)
	}
}
