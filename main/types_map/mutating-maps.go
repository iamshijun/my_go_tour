package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	//另外对于返回多个值的函数如果不需要除了第一个返回的值 都可以将后续的值忽略/省略
	// v := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
	//如果key不在 返回的value为 value类型的零值 (/ 0 / nil / ....)
	q := m["Question"]
	fmt.Println("The Question:", q)
}
