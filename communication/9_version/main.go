package main

/**
版本新增：
1.定时打印OnlineMap中的用户信息
*/
func main() {
	server := NewServer("127.0.0.1", 8888)
	server.StartServer()
}
