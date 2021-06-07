package main

import "fmt"

// 四种声明slice的方式
func main() {

	slice1 := []int{1, 2, 3}
	//var slice1 []int
	//var slice1 []int = make([]int, 10)
	//slice1 := make([]int, 10)

	// %v 打印所有详细信息
	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	// 判断一个slice是否分配了空间
	if slice1 == nil {
		fmt.Println("slice1 是一个空切片")
	} else {
		fmt.Println("slice1 不是一个空切片")
	}
}
