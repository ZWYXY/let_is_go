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

	// 让用户持有Server的引用
	server *Server
}

// 新建User
func NewUser(conn net.Conn, server *Server) *User {
	addr := conn.RemoteAddr().String()
	user := &User{
		Addr: addr,
		Name: addr,
		Conn: conn,
		C:    make(chan string),
		// 创建User时初始化server
		server: server,
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

// 上线功能
func (this *User) Online() {
	// 当前用户加入到Map中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Addr] = this
	this.server.mapLock.Unlock()
	// 广播消息，用户上线
	this.server.Broadcast(this, "用户上线")
}

// 下线功能
func (this *User) Offline() {
	// 当前用户从Map中移除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Addr)
	this.server.mapLock.Unlock()
	// 广播消息，用户下线
	this.server.Broadcast(this, "用户下线")
}

// 处理消息
func (this *User) DealMessage(msg string) {
	this.server.Broadcast(this, msg)
}
