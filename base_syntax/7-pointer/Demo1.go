package main

import "fmt"

/**
演示  Go语言中的指针
*/
func main() {
	a := 1
	changeValue(a)
	fmt.Printf("a = %d\n", a)

	fmt.Println("----------------")

	changeValue1(&a)
	fmt.Printf("a = %d\n", a)
}

func changeValue(b int) {
	b = 10
}

func changeValue1(b *int) {
	*b = 10
}
