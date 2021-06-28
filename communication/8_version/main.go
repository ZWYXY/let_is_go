package main

/**
版本新增：
1. 私聊功能
*/
func main() {
	server := NewServer("127.0.0.1", 8888)
	server.StartServer()
}
