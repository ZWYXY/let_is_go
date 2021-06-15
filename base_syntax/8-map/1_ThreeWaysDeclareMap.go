package main

import "fmt"

func main() {

	// 方式一：
	var myMap map[string]string
	fmt.Println("myMap = ", myMap)
	// 使用Map前需要先给Map分配空间
	myMap = make(map[string]string, 10)
	fmt.Println("myMap = ", myMap)

	// 方式二：
	myMap1 := make(map[string]string)
	//myMap1 := make(map[string]string, 10)
	myMap1["1"] = "go"
	myMap1["2"] = "c++"
	myMap1["3"] = "java"
	fmt.Println("myMap = ", myMap1)

	// 方式三：
	myMap2 := map[int]string{
		1: "go",
		2: "c++",
		3: "java",
	}
	fmt.Println("myMap = ", myMap2)

}
