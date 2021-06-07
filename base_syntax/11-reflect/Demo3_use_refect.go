package main

import (
	"fmt"
	"reflect"
)

// 反射机制的使用
func main() {

	animal := Animal{"green", 4, true}

	DoFilesAndMethod(animal)

	// 方法调用
	animal1 := &Animal{"black", 4., false}
	animal1Values := reflect.ValueOf(animal1)
	animal1Method := animal1Values.MethodByName("Eat")
	// 有参数就通过这里传参
	// paramList := []reflect.Value{}
	animal1Method.Call([]reflect.Value{})

}

type Animal struct {
	Color       string
	Legs        int
	ActiveInDay bool
}

func (this Animal) Eat() {
	fmt.Printf("animail eat foods...")
	fmt.Println()
}

func DoFilesAndMethod(input interface{}) {
	// 获取输入的type
	inputType := reflect.TypeOf(input)
	fmt.Println(inputType)

	// 获取输入的value
	inputValue := reflect.ValueOf(input)
	fmt.Println(inputValue)

	// 获取type中的字段名称、类型、值
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		// 注意如果Type中的字段不是public的（也就是不大写），这里会报错
		fmt.Printf("%s %v = %v\n", field.Name, field.Type, value)
	}

	// 反射获取type的方法信息
	for j := 0; j < inputType.NumMethod(); j++ {
		method := inputType.Method(j)
		fmt.Printf("%s, %v\n", method.Name, method.Type)
	}

}
