package main

import (
	"fmt"
	"time"
)

// 演示缓存channel的使用
func main() {

	/**
	创建一个带9缓存的channel
	当channel缓存用完后，向管道中输入数据的goroutine会陷入阻塞
	当channel为空时，从管道中取数据的goroutine会陷入阻塞
	*/
	c := make(chan int, 9)

	go func() {
		fmt.Println("goroutine 启动")
		for i := 0; i < 100; i++ {
			// 写数据到channel中
			c <- i
			fmt.Printf("i = %d, c 的 len = %d, cap = %d\n", i, len(c), cap(c))
		}
		defer fmt.Println("goroutine 结束")
	}()

	time.Sleep(3 * time.Second)

	// main goroutine中获取数据
	for i := 0; i < 2; i++ {
		// 从channel中读取数据
		x := <-c
		fmt.Printf("dataFromChannel = %d, c 的 len = %d, cap = %d\n", x, len(c), cap(c))
	}

	for {

	}
}
