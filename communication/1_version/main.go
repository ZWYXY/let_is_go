package main

/**
版本新增：
1. 基础服务搭建
*/
func main() {

	server := NewServer("127.0.0.1", 8080)
	server.StartServer()

}
