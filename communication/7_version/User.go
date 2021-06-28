package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Addr string
	Name string
	Conn net.Conn    // 连接
	C    chan string // 管道

	// 让用户持有Server的引用
	server *Server
}

// 新建User
func NewUser(conn net.Conn, server *Server) *User {
	addr := conn.RemoteAddr().String()

	user := &User{
		Addr: addr,
		Name: strings.Split(addr, ":")[1],
		Conn: conn,
		C:    make(chan string),
		// 创建User时初始化server
		server: server,
	}

	// 启用一个goroutine 监听当前User的Channel就是C
	go user.ListenMsg()

	return user
}

// 检测管道中的消息并发送到Conn中
func (this *User) ListenMsg() {
	fmt.Println(this.Addr + "的channel监听启动")
	for {
		// 这里监听User的管道，如果User的管道中没有信息会被阻塞，有信息就会读出来并回写给客户端
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
	fmt.Println(msg)
	if msg == "who" {
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			messageToSend := "用户地址：" + user.Addr + "，用户名：" + user.Name + "，在线。"
			this.SendMessage(messageToSend)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && "rename|" == msg[:7] {
		newName := strings.Split(msg, "|")[1]
		this.server.mapLock.Lock()
		user := this.server.OnlineMap[this.Addr]
		user.Name = newName
		this.server.mapLock.Unlock()
		this.SendMessage("成功修改昵称为：" + newName)
	} else {
		// 通过服务端广播，接收到的消息
		this.server.Broadcast(this, msg)
	}

}

func (this *User) SendMessage(msg string) {
	this.Conn.Write([]byte(msg))
}
