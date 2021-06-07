package main

import (
	"fmt"
	"time"
)

/**
演示，go中的对象的创建和使用
*/

type Person struct {
	Name   string
	Birth  time.Time
	gender bool
}

// 使用这种方式，只能读数据，不能改数据
/*
func (this Person) Show() {
	fmt.Printf("%v\n", this)
}

func (this Person) GetName() string {
	return this.Name
}

func (this Person) SetName(name string) {
	this.Name = name
}
*/

// 可读可改的写法

func (this *Person) Show() {
	fmt.Printf("%v\n", *this)
}

func (this *Person) GetName() string {
	return this.Name
}

func (this *Person) SetName(name string) {
	this.Name = name
}

func main() {

	person := Person{
		Name:   "123",
		Birth:  time.Now(),
		gender: true,
	}

	person.Show()
	person.SetName("法外狂徒")
	person.Show()
}
