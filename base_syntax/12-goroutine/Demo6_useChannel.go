package main

import "fmt"

// 演示select的使用
/*
单流程下（可以理解为一个main方法下创建了一个goroutine，有点不理解这里），
go只能监控一个channel的状态，如果有多个channel的状态需要监听，怎么办？因此出现了select
select {
case <- chan1:
case <- chan2:
case <- chan3:
default:
} 可以同一个流程下监控多个channel的状态
*/
func fibonacci(compute, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case compute <- x:
			x = y
			y = x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	compute, quit := make(chan int), make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-compute)
		}

		quit <- 0
	}()

	fibonacci(compute, quit)

}
