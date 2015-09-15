package main

import "fmt"

func main() {
	a := [6]int{2, 3, 5, 7, 11, 13} //这是一个slice不是数组(有固定大小)
	s := a[0:]
	sa := a[cap(a)-2 : cap(a)-1]
	//slice是指针, go中array可互相赋值,值传递即拷贝一份
	fmt.Printf("a => %v,len == %d, cap == %d\n", a, len(a), cap(a)) // [2 3 5 7 11 13] 6 6  (cap==len)
	fmt.Println("s == ", s)

	slice_rang1_4 := s[1:4] // [3 5 7] 3 5 (capacity=a容量 - 起始下标 = 6 - 1 = 5)
	fmt.Printf("s[1:4] => %v, len == %d, cap == %d\n",
		slice_rang1_4, len(slice_rang1_4), cap(slice_rang1_4))

	// 省略下标代表从 0 开始

	slice_rang0_3 := s[:3] // [2 3 5] 2 6  (capacity=a容量 - 起始下标 = 6 - 0 = 6)
	fmt.Printf("s[:3] => %v, len == %d, cap == %d\n",
		slice_rang0_3, len(slice_rang0_3), cap(slice_rang0_3))

	// 省略上标代表到 len(s) 结束
	slice_rang2_ := s[2:] // [5 7 11 13]
	fmt.Printf("s[2:] => %v, len == %d, cap == %d\n",
		slice_rang2_, len(slice_rang2_), cap(slice_rang2_))

	slice_rang1_1 := s[1:1]
	fmt.Printf("s[1:1] => %v, len == %d, cap == %d\n",
		slice_rang1_1, len(slice_rang1_1), cap(slice_rang1_1)) // []

	fmt.Println("length s[1:1] ==", len(s[1:1]))   //0
	fmt.Printf("type of slice == %T \n", (s[1:1])) // []int

	fmt.Printf("poninter address of s : %p\n", s)

	fmt.Printf("sa => %v,len == %d, cap == %d\n", sa, len(sa), cap(sa))
	sa = append(sa, 31)
	fmt.Printf("after append 31 into sa : %v, &sa : %p \n", sa, sa)

	sa = append(sa, 100, 101, 102)
	fmt.Printf("after append 100,101,102 into sa : %v, &sa : %p \n", sa, sa)

	fmt.Printf("then a is : %v, &a : %p \n", a, &a)
}
