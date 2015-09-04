package main

//!!应该叫 unbuffered-channel.go 汗
import (
	"fmt"
	"time"
)

//channel 是有类型的管道，可以用 channel 操作符 <- 对其发送或者接收值。
//（“箭头”就是数据流的方向。）
func sum(arr []int, c chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	c <- sum //将sum送入C
}

func main() {
	a := []int{7, 9, 3, 5, 1, -2}
	//和 map 与 slice 一样，channel 使用前>必须<创建：
	c := make(chan int)

	fmt.Printf("Type of channel : %T , len : %d, cap : %d\n", c, len(c), cap(c)) // chan int ,0 ,0

	go sum(a[:len(a)/2], c) // -> y
	go sum(a[len(a)/2:], c) // -> x
	//默认情况下，在另一端准备好之前，发送和接收都会阻塞。这使得 goroutine 可以在没有明确的锁或竞态变量的情况下进行同步。

	//所以对于当前的c (unbuffered channel)来说上述的两个goroutine(sum) 都会在 c<-num处阻塞,直到下面的receiver开始接收消息
	x, y := <-c, <-c
	//y := <-c

	//换种执行方式 如果让上述的两个goroutine先等待一段时间 让 x, y := <-c, <-c先执行 同样类似的情况发生在了receiver上,
	//这时候receiver会被阻塞 等待可用的消息返回 当前面的sum goroutine将数据传入channel中后 receiver会被唤醒继续执行

	//整体上可见 unbuffered channel 非常类似于 java的 SynchronousQueue (len always be 0),而对于有大小的channel就像是 LinkedBlockingQueue/ArrayBlockingQueue

	fmt.Println(x, y, x+y)
}
