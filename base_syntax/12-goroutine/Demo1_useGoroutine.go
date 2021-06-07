package main

import (
	"fmt"
	"time"
)

// 子routine
func newTask() {
	i := 0
	for {
		i++
		time.Sleep(1 * time.Second)
		fmt.Println("new Task print", i)
	}
}

// 主goroutine
func main() {

	// 开启 一个 goroutine 去执行newTask()
	go newTask()

	i := 0

	for {
		i++
		time.Sleep(1 * time.Second)
		fmt.Println("main print", i)
	}
}
