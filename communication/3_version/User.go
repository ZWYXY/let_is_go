package main

import (
	"fmt"
	"net"
)

type User struct {
	Addr string
	Name string
	Conn net.Conn
	C    chan string
}

// 新建User
func NewUser(conn net.Conn) *User {
	addr := conn.RemoteAddr().String()
	user := &User{
		Addr: addr,
		Name: addr,
		Conn: conn,
		C:    make(chan string),
	}

	// 启动当前user Channel消息的goroutine
	go user.ListenMsg()

	return user
}

// 检测管道中的消息并发送到Conn中
func (this *User) ListenMsg() {
	for {
		// 这里监听User的管道，如果User的管道中没有信息会被阻塞，有信息就会读出来并回写给客户端
		fmt.Println(this.Addr + "的channel监听启动")
		msg := <-this.C // 无法从管道中获取数据时，会陷入阻塞
		this.Conn.Write([]byte(msg + "\n"))
	}
}
