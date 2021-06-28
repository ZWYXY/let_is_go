package main

/**
版本新增：
1. 对用户的一些方法进行封装
*/
func main() {
	server := NewServer("127.0.0.1", 8888)
	server.StartServer()
}
