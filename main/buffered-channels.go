package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 2) //第二个参数不给定时 为unbuffered channel
	fmt.Printf("after add val 1, len : %d, cap : %d\n", len(c), cap(c))

	c <- 2
	fmt.Printf("after add val 2, len : %d, cap : %d\n", len(c), cap(c))
	//channel也存在着 len, cap 当len=cap=0表示缓冲区无大小限制
	//len为当前缓冲区含有数据的大小,cap为channel容量

	go func(c chan int, seconds int) {
		time.Sleep(time.Second)
		fmt.Println("Get ", <-c)
	}(c, 2)

	//向缓冲 channel 发送数据的时候，只有在缓冲区满的时候才会阻塞。当缓冲区清空的时候接受阻塞

	fmt.Println("add third val into channel", time.Now())
	c <- 3 //we block here waiting the space
	fmt.Println("after add third val into", time.Now())

	fmt.Println(<-c)
	fmt.Println(<-c)
}
