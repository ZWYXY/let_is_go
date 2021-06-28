package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	go count()

	// go 语言字符串切割
	s := "127.0.0.1:8080"
	split := strings.Split(s, ":")
	for i, n := range split {
		fmt.Printf("i = %d, n = %s\n", i, n)
	}

	// 字节数组
	myByteArray := []byte{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println(myByteArray)
	fmt.Println(myByteArray[0]) // slice(动态数组)从0开始
	fmt.Println(myByteArray[1])
	fmt.Println(myByteArray[4])
	fmt.Println(myByteArray[:4]) // 从0开始截取，[0,4)   前开后闭，懂我意思吧
	fmt.Println(myByteArray[4:]) // 从4开始截取，[4,end) 前开后闭，懂我意思吧

	select {}

}

func count() {
	for {
		fmt.Println("hit")
		time.Sleep(1 * time.Second)
	}
}
