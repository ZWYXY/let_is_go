package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

var ServerIp string
var ServerPort int

/**
1. 菜单显示
*/
func main() {

	// 命令行解析
	flag.Parse()

	client := NewClient(ServerIp, ServerPort)
	if client == nil {
		fmt.Println(">>>>>> 连接服务器失败")
		return
	}

	// 在连接服务器成功后，开启go程专门处理服务器响应
	go client.DealResponse()

	fmt.Println(">>>>>> 连接服务器成功")

	client.Run()

	//time.Sleep(10 * 60 * time.Second)
}

// 定义一个客户端类型（结构体）
type MyClient struct {
	// IP
	ServerIp string
	// 端口
	ServerPort int
	// 名称
	ServerName string
	// 连接
	ServerConn net.Conn
	// flag
	flag int
}

// 用来解析命令行参数
func init() {
	flag.StringVar(&ServerIp, "ip", "127.0.0.1", "设置服务器Ip默认是（127.0.0.1）")
	flag.IntVar(&ServerPort, "port", 88888, "设置服务器port默认是（88888）")
}

// 新建一个Client对象
func NewClient(ip string, port int) *MyClient {

	client := &MyClient{
		ServerIp:   ip,
		ServerPort: port,
		flag:       999,
	}

	dial, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Println(err)
	}

	client.ServerConn = dial

	return client
}

// 菜单方法
func (client *MyClient) menu() bool {
	var flag int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.修改昵称")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println("请输入正确的模式编号")
		return false
	}
}

// 运行菜单方法
func (client *MyClient) Run() {
	for client.flag != 0 {
		for !client.menu() {
		}
		switch client.flag {
		case 1:
			fmt.Println("进入公聊模式")
		case 2:
			fmt.Println("进入私聊模式")
		case 3:
			// 更新用户名
			client.UpdateUsername()
		}
	}
}

// 修改用户名
func (client *MyClient) UpdateUsername() {
	fmt.Println("请输入新的用户名 >>>>")

	var newName string
	fmt.Scanln(&newName) // 获取用户输入的新用户名

	// 发送指令
	cmd := "rename|" + newName
	client.ServerConn.Write([]byte(cmd))
}

func (client *MyClient) DealResponse() {
	io.Copy(os.Stdout, client.ServerConn)
}
