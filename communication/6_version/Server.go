package main

import (
	"fmt" //
	"io"
	"net"  // 网络通信功能相关包
	"sync" // 同步功能相关api包
)

// Server 结构体
type Server struct {
	Ip        string           // ip
	Port      int              // 端口
	OnlineMap map[string]*User // 在线用户表
	mapLock   sync.RWMutex     // 读写锁
	Message   chan string      // 管道
}

// NewServer 创建Server对象
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// StartServer Server的启动方法
func (this *Server) StartServer() {
	// 调用net包下的listen方法开启一个监听器
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	defer listener.Close()

	// 开启一个goroutine，监听Message管道
	go this.ListenMessage()

	// 循环监听是否如果有连接建立
	for {
		conn, err := listener.Accept() // 如果没有客户端接入，这里会一直阻塞，下面的代码就无法执行
		if err != nil {
			fmt.Println("net.Listen err:", err)
			continue
		}
		// 有连接建立，开启goroutine处理业务
		go this.BusinessHandler(conn)
	}
}

// BusinessHandler 业务处理类
func (this *Server) BusinessHandler(conn net.Conn) {
	// 创建User
	user := NewUser(conn, this)

	// invoke用户上线方法
	user.Online()

	// 从管道中读取数据，并交给用户处理
	go func() {
		for {
			bytes := make([]byte, 1024)
			// 从连接（Conn）中读取用户消息 读不到信息，这里会阻塞
			n, err := conn.Read(bytes)
			fmt.Println(n)
			if n == 0 {
				// 用户下线
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println(err)
				return
			}

			// 提取用户消息 去除 '\n'
			msg := string(bytes[:n])
			//msg := string(bytes)

			// 用户发送消息
			user.DealMessage(msg)
		}
	}()

	// 阻塞处理
	select {}
}

// 回写消息到管道
func (this *Server) Broadcast(user *User, msg string) {
	msgToBeSend := "\n[ " + user.Addr + " ] " + user.Name + ":" + msg
	// 回写消息到Message管道
	this.Message <- msgToBeSend
}

// 从管道中取出消息并回写给客户端
func (this *Server) ListenMessage() {
	for {
		// 监听管道中的信息，回写给客户端
		msg := <-this.Message

		this.mapLock.Lock()
		for _, user := range this.OnlineMap {
			// 这里向User的管道里面写入数据，会触发User对象的ListenMsg方法
			user.C <- msg
		}
		this.mapLock.Unlock()
	}
}
