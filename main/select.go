package main

import "fmt"

// c chan<- int, quit <-chan int /* c is only a receiver , quit is a comsumer*/
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: //当下述的goroutine执行完循环n次取channel数据后 n+1次的时候当前case中
			//c <- x语句 会被阻塞 因为当前没有任何的recevier , consumer会被阻塞 进而判断 <-quit语句
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

/*
	select 语句使得一个 goroutine 在多个通讯操作上等待。
	select 会阻塞，直到条件分支中的某个可以继续执行，这时就会执行那个条件分支。
	当多个都准备好的时候，会随机选择一个。
*/

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
