package main

/**
版本新增：
1. 查询在线用户功能
*/
func main() {
	server := NewServer("127.0.0.1", 8888)
	server.StartServer()
}
