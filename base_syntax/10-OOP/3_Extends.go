package main

import "fmt"

/*
	演示继承
*/

type Human struct {
	name   string
	gender string
}

// Human name属性的get set 方法
func (this *Human) SetName(name string) {
	this.name = name
}

func (this Human) getName() string {
	return this.name
}

func (this Human) eat() {
	fmt.Printf("Human's eat method\n")
}

func (this Human) walk() {
	fmt.Printf("Human's walk method\n")
}

type SuperMan struct {
	Human // 继承了父类Human
	age   string
}

func (this SuperMan) fly() {
	fmt.Printf("Superman can fly\n")
}

func (this SuperMan) print() {
	fmt.Printf("SuperMan %v\n", this)
}

func main() {

	aSuperMan := SuperMan{Human{"123", "男"}, "Never old"}
	aSuperMan.eat()
	aSuperMan.fly()
	aSuperMan.walk()
	aSuperMan.print()

	var oneSuperMan SuperMan
	oneSuperMan.age = "18"
	oneSuperMan.gender = "男"
	oneSuperMan.name = "one Super man "
	oneSuperMan.print()
}
