package main

import "fmt"

type Count int

func (count *Count) Increment()  { *count++ } // Count类型的方法
func (count *Count) Decrement()  { *count-- }
func (count Count) IsZero() bool { return count == 0 }

type Part struct { //基于结构体创建自定义类型Part
	stat  string
	Count //匿名字段 ps 何用
}

func (part Part) IsZero() bool { // 覆盖??了匿名字段Count的IsZero()方法
	return part.Count.IsZero() && part.stat == "" //调用匿名字段的方法
}

func (part Part) String() string { //实现Stringer interface
	return fmt.Sprintf("<<%s, %d>>", part.stat, part.Count)
}

func main() {
	var i Count = -1
	fmt.Printf("Start \"Count\" test:\nOrigin value of count: %d\n", i)

	i.Increment() //(&i).Increment()

	/*在以上定义的类型Count中，*Count方法集是Increment(), Decrement()和IsZero()，Count的值的方法集是IsZero()。
	但是因为Count类型的是可寻址的，所以我们可以使用Count的值调用全部的方法*/

	fmt.Printf("Value of count after increment: %d\n", i)
	fmt.Printf("Count is zero t/f? : %t\n\n", i.IsZero())
	fmt.Println("Start: \"Part\" test:")

	part := Part{"232", 0}
	fmt.Printf("Part: %v\n", part)                                      //String() method invoked
	fmt.Printf("Part is zero t/f? : %t\n", part.IsZero())               //假设part没有自己的IsZero方法会找 匿名字段Count的IsZero方法,注释掉看看!
	fmt.Printf("Count in Part is zero t/f?: %t\n", part.Count.IsZero()) // 尽管覆盖了匿名字段的方法，单还是可以访问

}
