package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("-----foo1-----")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	c := 100 + b

	return c
}

// 多返回值匿名的
func foo2(a string, b int) (int, string) {
	fmt.Println("-----foo2-----")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	c := 100 + b

	return c, a
}

// 多返回值，有形参名称的
func foo3(a string, b int) (res1 int, res2 string) {
	fmt.Println("-----foo3-----")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// 形参默认值打印
	fmt.Println("res1 =", res1)
	fmt.Println("res2 =", res2)

	res1 = b + 100 // res1 res2 范围，整个函数内有效
	res2 = a + ""

	return
}

func foo4(a string, b int) (res1, res2 int) { // res1 和 res2 返回值类型一样的
	fmt.Println("-----foo4-----")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	res1 = b + 100 // res1 res2 范围，整个函数内有效
	res2 = b + 200

	return
}

func main() {
	c := foo1("11", 123)
	fmt.Println("c =", c)

	res1, res2 := foo2("12", 1000)
	fmt.Printf("res1 = %d, res2 = %s\n", res1, res2)

	res3, res4 := foo3("34", 10000)
	fmt.Printf("res3 = %d, res4 = %s\n", res3, res4)

	res5, res6 := foo4("56", 100000)
	fmt.Printf("res5 = %d, res6 = %d\n", res5, res6)

}
