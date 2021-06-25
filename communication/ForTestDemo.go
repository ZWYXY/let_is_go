package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	go count()

	// go 语言字符串切割
	s := "127.0.0.1:8080"
	split := strings.Split(s, ":")
	for i, n := range split {
		fmt.Printf("i = %d, n = %s", i, n)
	}

	select {}

}

func count() {
	for {
		fmt.Println("hit")
		time.Sleep(1 * time.Second)
	}
}
