package main

import "fmt"

// 全局变量只能使用 方式1，2，3  := 不能使用

// 演示变量的四种声明方式
func main() {

	// 方法一： 只声明不赋值
	var a int
	fmt.Printf("a = %d , a.type = %T\n", a, a)

	// 方法二：声明并赋值
	var b int = 100
	fmt.Printf("b = %d , b.type = %T\n", b, b)

	// 方法三：去掉类型说明，编译器自动推断出类型，因此需要赋值
	var c = 100
	fmt.Printf("c = %d , c.type = %T\n", c, c)

	// 方法四：:=
	d := 100
	fmt.Printf("d = %d , d.type = %T\n", d, d)

	// 多变量声明
	var e, f int = 100, 20
	fmt.Printf("e = %d , e.type = %T\n", e, e)
	fmt.Printf("f = %d , f.type = %T\n", f, f)

	var g, h = "100", 100
	fmt.Printf("g = %s , g.type = %T\n", g, g)
	fmt.Printf("h = %d , h.type = %T\n", h, h)

	var (
		i = 100
		j = "100"
		k = false
	)
	fmt.Printf("i = %d , i.type = %T\n", i, i)
	fmt.Printf("j = %s , j.type = %T\n", j, j)
	fmt.Printf("k = %t , k.type = %T\n", k, k)

	l, m := "100", 100
	fmt.Printf("g = %s , g.type = %T\n", l, l)
	fmt.Printf("h = %d , h.type = %T\n", m, m)

}
