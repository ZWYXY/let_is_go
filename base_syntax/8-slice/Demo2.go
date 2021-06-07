package main

import "fmt"

// 因为Demo1中方法传数组不方便（只能传固定的长度的数组进来，且是值传递），因此有了slice（切片） 也就是动态数组
func main() {

	array1 := []int{1, 2, 3, 4}          // 动态数组，切片 slice
	array2 := []int{1, 2, 3, 4, 6, 7, 8} // 动态数组，切片 slice

	// 演示动态数组
	dynamicArray(array1)

	fmt.Println("------------------------")

	//　go中 _ 表示匿名（联想匿名引用包是怎么引用的），匿名了当然值也就拿不到了
	for _, value := range array1 {
		fmt.Printf("value = %d\n", value)
	}
	fmt.Println("------------------------")

	dynamicArray(array2)
	fmt.Println("------------------------")
	for _, value := range array1 {
		fmt.Printf("value = %d\n", value)
	}
}

// 可以传任意长度的数组进来，并且对其进行修改会改变原数组的值（引用传递，就是传地址）
func dynamicArray(array1 []int) {
	for index, value := range array1 {
		fmt.Printf("index = %d, value = %d\n", index, value)
	}

	array1[0] = 1000000000000
}
