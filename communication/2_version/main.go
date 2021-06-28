package main

/**
版本新增：
1. 用户上线，消息广播
*/
func main() {
	server := NewServer("127.0.0.1", 8080)
	server.StartServer()
}
