package main

import "fmt"

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *Tree, ch chan int) {
	ch <- t.Value
	go Walk(t.Left, ch)
	go Walk(t.Right, ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *Tree) bool {
	c1, c2 := make(chan int), make(chan int)
	a1, a2 := make([]int, 10), make([]int, 10)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for {
		select {
		case v1 := <-c1:
			a1 = append(a1, v1)
		case v2 := <-c2:
			a1 = append(a2, v2)
		default:
			break
		}
	}

	if len(a1) != len(a2) {
		return false
	} else {
		//sort a1, a2
		sort(a1)
		sort(a2)

		//compare a1 and a2
		for i, la := 0, len(a1); i < la; i++ {
			if a1[i] != a2[i] {
				return false
			}
		}

		return true
	}

}

func sort(a []int) {
	for i := len(a) - 1; i >= 1; i-- { //确定右限
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func main() {
	a := []int{6, 2, 4, 5, 1, 8, 0, 9, 7, 3}
	fmt.Println(a)

	sort(a)
	fmt.Println(a)

}
