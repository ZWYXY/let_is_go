package main

import (
	"fmt"
	"time"
)

// 演示缓存channel的使用
func main() {
	defer fmt.Println("main finished")

	/**
	创建一个带9缓存的channel
	当channel为空时，从管道中取数据的goroutine会陷入阻塞
	当channel缓存用完后，向管道中输入数据的goroutine会陷入阻塞
	*/
	c := make(chan int, 9)

	go func() {
		defer fmt.Println("goroutine finished")
		for i := 0; i < 10; i++ {
			c <- i
			fmt.Printf("i = %d, c 的 len = %d, cap = %d\n", i, len(c), cap(c))
		}
	}()

	time.Sleep(3 * time.Second)

	// main goroutine中获取数据
	for i := 0; i < 10; i++ {
		x := <-c
		fmt.Printf("x = %d, c 的 len = %d, cap = %d\n", x, len(c), cap(c))
	}

}
