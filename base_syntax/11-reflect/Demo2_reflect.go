package main

import "fmt"

type Reader interface {
	read()
}

type Writer interface {
	write()
}

type Book struct {
}

func (this *Book) read() {
	fmt.Println("read a book")
}

func (this *Book) write() {
	fmt.Println("write a book")
}

func main() {

	//pair<type:Book, value:oneBook的地址>
	//oneBook := &Book{}
	var oneBook = &Book{}

	//pair<type:, value:>
	var r Reader
	//pair<type:Book, value:oneBook的地址>
	r = oneBook
	r.read()

	//pair<type:, value:>
	var w Writer
	//pair<type:Book, value:oneBook的地址>
	w = r.(Writer) // 这里为什么可以断言成功？因为 w 和 r的 type 一致
	w.write()

}
