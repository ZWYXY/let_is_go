package main

import "fmt"

/**
演示万能类型与断言机制
*/
func main() {

	myFunc("111")
	myFunc(111)
	myFunc(false)
	myFunc(Bank{"go", "let's go"})
}

func myFunc(arg interface{}) {
	fmt.Printf("my Func is called!")

	// 断言机制 只允许使用在interface{}中
	value, ok := arg.(string)
	if ok {
		fmt.Printf("%s is string type \n", value)
	} else {
		fmt.Println("arg is not string type")
	}
}

type Bank struct {
	address string
	name    string
}
