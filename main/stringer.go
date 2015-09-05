package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	//Sprint donot print out but return the string output
	return fmt.Sprintf("%q (%v years)", p.Name, p.Age)
}

//Stringer 是一个可以用字符串描述自己的类型。`fmt`包 （还有许多其他包）使用这个来进行输出。
func main() {
	a := Person{"Authur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	//just like java Object.toString() , class can override this method
	fmt.Println(a, z)
}
