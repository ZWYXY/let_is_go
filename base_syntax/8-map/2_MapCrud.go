package main

import "fmt"

/**
演示map crud
*/
func main() {

	myMap := map[string]string{
		"China": "Beijing",
		"USA":   "NewYork",
		"Japan": "Tokyo",
	}

	// go中 对于map传递的是引用，也就是说，在这里的changeMyMap中修改map的值，会改边main方法中myMap
	changeMyMap(myMap)

	// print in main
	fmt.Println()
	fmt.Printf("print in main ----------------\n")
	for key, value := range myMap {
		fmt.Printf("key = %s, value = %s\n", key, value)
	}
}

func changeMyMap(myMap map[string]string) {

	// 遍历
	for key, value := range myMap {
		fmt.Printf("key = %s, value = %s\n", key, value)
	}
	fmt.Printf("add new ----------------\n")
	// 新增
	myMap["England"] = "London"
	for key, value := range myMap {
		fmt.Printf("key = %s, value = %s\n", key, value)
	}
	fmt.Printf("del ----------------\n")
	// 删除
	delete(myMap, "Japan")
	for key, value := range myMap {
		fmt.Printf("key = %s, value = %s\n", key, value)
	}

	fmt.Printf("update ----------------\n")
	// 修改
	myMap["USA"] = "DC"
	for key, value := range myMap {
		fmt.Printf("key = %s, value = %s\n", key, value)
	}

	// 查询
	fmt.Printf("slect ----------------\n")
	fmt.Printf("value = %s\n", myMap["China"])
}
