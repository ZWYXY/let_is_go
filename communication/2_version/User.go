package main

import "net"

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
	go user.ListenMsg()
	return user
}

// 检测管道中的消息并发送到Conn中
func (this *User) ListenMsg() {
	for {
		msg := <-this.C
		this.Conn.Write([]byte(msg + "\n"))
	}
}
