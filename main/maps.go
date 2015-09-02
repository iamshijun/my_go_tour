package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// string(key type) Vertex(value type)
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	//v := Vertex{40.68433, -74.39967} //如果像上面那样换行的话 最后需要一个 "," why?
	v := Vertex{Lat: 1.09}
	fmt.Println(v)
}
