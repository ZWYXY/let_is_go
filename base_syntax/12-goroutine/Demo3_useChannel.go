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
		// 方式一： 接收并将其丢弃
		<- c
	*/

	/*
		// 方式二：接受值并赋值
		x := <- c
	*/
	// 方式三： 接收并检查通道是否已关闭或者是否为空 true 表示channel没有关闭， false表示channel已经关闭
	x, ok := <-c
	fmt.Println(x)
	fmt.Println(ok)
	fmt.Println("end main goroutine")
}
