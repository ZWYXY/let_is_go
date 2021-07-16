package main

import "fmt"

/**
Go语言的面向对象编程，这里定义了一个Person对象
*/

type Person struct {
	name string
}

/**
下面这两个是特殊的function叫做method一般是用来操作对象的
*/

func (receiver Person) GetName() (name string) {
	return receiver.name
}

/**
注意这里的Person前有个*表示这里传入的是调用方的指针，因此可以操作原对象
如果没有加 * 这个setName方法是不会生效的
*/

func (receiver *Person) SetName(name string) {
	receiver.name = name
}

func main() {
	p1 := &Person{name: "hhh"}
	fmt.Printf("person name is %s\n", p1.GetName())
	p1.SetName("我是谁？")
	fmt.Printf("person name is %s\n", p1.GetName())
}
