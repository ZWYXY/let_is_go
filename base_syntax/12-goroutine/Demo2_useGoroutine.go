package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	// 匿名函数的使用
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			// return 只能退出此匿名函数，进而退出go程
			runtime.Goexit() // 可以直接退出go程
			fmt.Println("B")
		}() // 这里需要调用一下定义的匿名函数

		fmt.Println("A")
	}()

	fmt.Println("main")

	// 无法通过接收返回值，获取执行结果，需要通过channel
	go func(a int, b int) bool {
		fmt.Printf("a = %d, b = %d", a, b)
		return (a - b) > 0
	}(10, 20)

	for {
		time.Sleep(1 * time.Second)
	}
}
