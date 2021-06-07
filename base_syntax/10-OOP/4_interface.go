package main

import "fmt"

/*
	演示 interface
*/

type AnimalIF interface { // interface 本质是一个指针
	Eat()
	Bark()
	Sleep()
}

type Cat struct {
	color string
}

func (this *Cat) Eat() {
	fmt.Printf(this.color + "Cat eat food!")
	fmt.Println()
}

func (this *Cat) Bark() {
	fmt.Printf(this.color + "Cat barked!")
	fmt.Println()
}

func (this *Cat) Sleep() {
	fmt.Printf(this.color + "Cat Sleeped!")
	fmt.Println()
}

type Dog struct {
	kind string
}

func (this *Dog) Eat() {
	fmt.Printf(this.kind + "Dog eat food!")
	fmt.Println()
}

func (this *Dog) Bark() {
	fmt.Printf(this.kind + "Dog barked!")
	fmt.Println()
}

func (this *Dog) Sleep() {
	fmt.Printf(this.kind + "Dog Sleeped!")
	fmt.Println()
}

func print(animalIF AnimalIF) {
	animalIF.Eat()
	animalIF.Bark()
	animalIF.Sleep()
}

func main() {
	oneCat := Cat{"Green"}
	oneDog := Dog{"牧羊犬"}

	print(&oneCat)
	print(&oneDog)
}
