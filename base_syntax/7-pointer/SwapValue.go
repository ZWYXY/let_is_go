package main

import "fmt"

/**
演示 交换值
*/
func main() {
	a := 1
	b := 2
	swapValue(&a, &b)
	fmt.Printf("a = %d, b = %d\n", a, b)
}

// 交换值
func swapValue(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}
