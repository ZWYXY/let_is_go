package main

import "fmt"

// defer的特性和执行顺序
func main() {
	// defer 后接语句 表示这个语句在方法结束后再执行 也就是return后执行
	// defer 使用栈，先进后出，在这里表现为，先定义后执行
	defer fmt.Println("defer print1 .....")
	defer fmt.Println("defer print2 .....")
	defer fmt.Println("defer print3 .....")

	fmt.Println("main print1 .....")
	fmt.Println("main print2 .....")

	// 代码执行结果：
	// main print1 .....
	// main print2 .....
	// defer print3 .....
	// defer print2 .....
	// defer print1 .....
}
