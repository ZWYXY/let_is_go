package main

import "fmt"

// 证明 return比defer先执行
func returnTask() int {
	fmt.Println("returnTask print .....")
	return 0
}

func deferTask() int {
	defer fmt.Println("deferTask print .....")
	return returnTask()
}

func main() {
	deferTask()
}
