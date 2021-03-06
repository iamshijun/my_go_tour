package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, 0.12, "no1"

	fmt.Println(i, j, k, c, python, java)

	for i := 0; i < k; i++ { //here we cannot use "var i = 0"
		fmt.Printf("% d", i)
	}

	/*
		在函数中, := 简洁赋值语句在明确类型的地方,可以用于替代 var 定义。
		函数外的每个语句都必须以关键字开始（ var 、 func 、等等）, := 结构不能使用在函数外。
	*/
}
