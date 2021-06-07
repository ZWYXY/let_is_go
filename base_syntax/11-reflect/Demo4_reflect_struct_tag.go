package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string `info:"用户名"  doc:"username"`
	Password string `info:"密 码"   doc:"password"`
}

func printStructTag(arg interface{}) {

	argType := reflect.TypeOf(arg)

	for i := 0; i < argType.NumField(); i++ {
		info := argType.Field(i).Tag.Get("info")
		doc := argType.Field(i).Tag.Get("doc")
		fmt.Printf("info = %s doc = %s\n", info, doc)
	}
}

func main() {
	var user User

	printStructTag(user)

	user.alterValue()

	fmt.Println(user)

}

func (this *User) alterValue() {
	this.Username = "hello"
	this.Password = "world"
}
