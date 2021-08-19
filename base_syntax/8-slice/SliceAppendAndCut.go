package main

import "fmt"

/**
slice 允许动态扩容
slice 需要注意的是，对slice的操作，一般都是基于引用的，也就是说，追加，截取，后的结果，再进行操作会影响到原数组
*/
func main() {
	slice1 := make([]int, 3) // 没有指定扩容多少，就按照初始化时指定的元素数量扩容 也就是3
	//slice1 := make([]int, 3, 5) //代表着 初始化前三个元素，但是分配5个元素的空间，扩容时每次扩容5个元素的空间
	fmt.Printf("len = %d, cap = %d, slice1 = %v\n", len(slice1), cap(slice1), slice1)

	// 追加元素
	slice1 = append(slice1, 10)
	fmt.Printf("len = %d, cap = %d, slice1 = %v\n", len(slice1), cap(slice1), slice1)

	// 截取元素
	slice2 := slice1[1:3]
	fmt.Printf("len = %d, cap = %d, slice2 = %v\n", len(slice2), cap(slice2), slice2)

	// 尝试修改截取后的数组，看看原数组是否有变化
	slice2[0] = 5
	fmt.Printf("原数组\n")
	fmt.Printf("len = %d, cap = %d, slice1 = %v\n", len(slice1), cap(slice1), slice1)
	fmt.Printf("从原数组中截取出来的数组\n")
	fmt.Printf("len = %d, cap = %d, slice2 = %v\n", len(slice2), cap(slice2), slice2)

	// 修改了截取后的数组发现，原数组中相同位置的元素也发生了改变，所以需要注意

	// 这里提供一个拷贝数组的方式
	slice3 := make([]int, 3, 3)
	copy(slice3, slice1)
	for index, value := range slice3 {
		fmt.Printf("index = %d, value = %d\n", index, value)
	}

	slice4 := make([]int, 0, 2)
	slice4 = append(slice4, 1)
	slice4 = append(slice4, 2)
	slice4 = append(slice4, 3)
	for index, value := range slice4 {
		fmt.Printf("index = %d, value = %d, cap = %d\n", index, value, cap(slice4))
	}
}
