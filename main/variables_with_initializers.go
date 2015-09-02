package main

import "fmt"

var i, j int = 1, 2

func main() {
	//如果初始化是使用表达式expression，则可以省略类型；变量从初始值中获得类型。
	var c, python, java, odd = true, false, "no!", 1
	fmt.Println(i, j, c, python, java, odd)

	//什么才算是表达式!!!
}
