package main

import "fmt"

// 演示select的使用

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
