package main

/**
版本新增：
1. 修改用户昵称功能 用户需要发送 rename|nickname
*/
func main() {
	server := NewServer("127.0.0.1", 8888)
	server.StartServer()
}
