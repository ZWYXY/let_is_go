package main

import "fmt"

// 结束一个goroutine
func main() {
	defer fmt.Printf("main finished!")

	//  开启channel
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}

		// 关闭channel
		// 关闭channel后还可以继续接收数据，但是发送数据
		close(c)
	}()

	//for i := 0; i < 10; i++ {
	//	x := <- c
	//	fmt.Printf("x = %d", x)
	//}

	/*	for {
		if data, ok := <- c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}*/
	// 循环从channel中拿数据，可以这么写
	for data := range c {
		fmt.Println(data)
	}

}
