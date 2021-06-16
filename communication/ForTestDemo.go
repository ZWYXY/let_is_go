package main

import (
	"fmt"
	"time"
)

func main() {

	go count()

	select {}

}

func count() {
	for {
		fmt.Println("hit")
		time.Sleep(1 * time.Second)
	}
}
