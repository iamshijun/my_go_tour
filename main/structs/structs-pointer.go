package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}

	//pv := &v
	var pv *Vertex = &v
	pv.X = 1e9 //1000000000
	//pointer to struct,property access we donnot need "->", just use "."
	fmt.Println(v)
}
