package main

import "fmt"

/**
演示如何使用结构体
Go中 方法首字母大写，属性首字母大写表示 共有，可以在包外访问到，首字母小写表示只有包内可以访问
*/

type Book struct {
	title string
	auth  string
}

func changeBook(book Book) {
	// 传递一个book副本
	book.title = "Get go!"
}

func changeBook1(book *Book) {
	book.title = "Get go!"
}

func main() {

	book := Book{}

	book.title = "Go从入门到精通"
	book.auth = "aceld"
	fmt.Printf("%v\n", book)

	fmt.Println("after change book value ---------------")

	changeBook(book)
	fmt.Printf("使用值传递后的book ----- %v\n", book)

	changeBook1(&book)
	fmt.Printf("使用地址（引用）传递后的book ----- %v\n", book)
}
