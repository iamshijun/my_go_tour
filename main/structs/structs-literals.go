package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // 类型为 Vertex , shi_ps constructor use index
	v2 = Vertex{X: 1}  // Y:0 被省略 , shi_ps constructor use type like map
	v3 = Vertex{}      // X:0 和 Y:0
	p  = &Vertex{1, 2} // 类型为 *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
	//{1 2} &{1 2} {1 0} {0 0}
	//方式指针都会输出一个 &前缀 然后是实际结构体的内容
}
