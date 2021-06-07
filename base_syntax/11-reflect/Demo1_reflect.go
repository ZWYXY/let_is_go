package main

import "fmt"

/*
	<p>
		go 中变量 有两个部分 一个是type, 另一个是value, 称为一个pair  pair<type:string value:"123">
		其中type 又分为static type 和 concrete type
	</p>
	<p>
		static type : 就是编码时可见的，常见的如：int string float32比如我们定义一个 var a int = 10 这个int类型 就是static type
		concrete type : 是runtime type系统看见的类型，也就是interface{}类型一个指针指向的concrete type
	</p>
	<p>一个interface{}类型的变量包含了2个指针，一个指针指向值的类型【对应concrete type】，另外一个指针指向实际的值【对应value】。</p>
*/
// 本demo演示这个pair是存在的
func main() {

	var a string
	// pair<statictype:string, value:"HelloWorld>
	a = "HelloWorld"

	// pair<concretetype:, value:>
	var allType interface{}

	// pair<concretetype:string, value:"HelloWorld">
	allType = a // 赋值时，会把type 和 value 带过去，保持不变 也就是pair不变，会一直带过去

	value, _ := allType.(string) // 这里能够断言是因为，a的type 和 allType的 concrete type 都是string
	fmt.Println(value)
}
