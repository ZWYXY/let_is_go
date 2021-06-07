package main

import "fmt"

// 线程间间通信 channel
func main() {

	// 定义channel
	c := make(chan int) // 创建无缓存channel

	go func() {
		fmt.Println("goroutine 1")
		c <- 1000 // 发送value 到channel
	}()

	// channel 接收值的三种方式
	/*
		// 接收并将其丢弃
		<- c
	*/

	/*
		// 接受值
		x := <- c
	*/
	x, ok := <-c //接收并检查通道是否已关闭或者是否为空 true 表示channel没有关闭， false表示channel已经关闭
	fmt.Println(x)
	fmt.Println(ok)
	fmt.Println("end main goroutine")
}
