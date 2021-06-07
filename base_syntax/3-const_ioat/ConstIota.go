package main

import "fmt"

// 用const来定义枚举类型
const (
	/*
		iota关键字，第一行的iota的默认值是0 后面每行都会iota累加1 后赋值给指定的变量
		iota只能和const配合使用，也就是只能写在const () 中
	*/
	BEIJING = iota
	SHANGHAI
	SHENZHEN
	GUANGZHOU
)

// 演示常量的使用
func main() {

	// 声明一个常量， 什么是常量？只读属性，不允许被修改
	const length int = 10
	fmt.Printf("length = %d\n", length)
	/*
		//尝试修改被const变量分配值
		length = 11;
		cannot assign to length (declared const)
	*/

	fmt.Printf("BEIJING = %d\n", BEIJING)
	fmt.Printf("BEIJING = %d\n", SHANGHAI)
	fmt.Printf("BEIJING = %d\n", SHENZHEN)
	fmt.Printf("BEIJING = %d\n", GUANGZHOU)

}
