package main

import "fmt"

// 演示数组的定义和遍历
func main() {

	// 方式一：
	var array1 [10]int
	// 方式二：
	array2 := [10]int{1, 2, 3, 4}

	// 遍历方式一
	for i := 0; i < len(array1); i++ {
		fmt.Printf("index = %d, value = %d\n", i, array1[i])
	}

	fmt.Println("---------------------")

	// 遍历方式二
	for index, value := range array2 {
		fmt.Printf("index = %d, value = %d\n", index, value)
	}

	fmt.Println("---------------------")

	// 调用
	passArray(array1)

	fmt.Println("---------------------")
	// 检查原数组
	for index, value := range array1 {
		fmt.Printf("index = %d, value = %d\n", index, value)
	}
}

// 方法传数组
// 需要注意的是，这里的数组，定义了长度是10，就只能传10长度的数组进来，且不会改变原来数组的值（因为这里的array1 是原数组的copy，值传递嘛）
func passArray(array1 [10]int) {

	for index, value := range array1 {
		fmt.Printf("index = %d, value = %d\n", index, value)
	}

	array1[0] = 1000
}
