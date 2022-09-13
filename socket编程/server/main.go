package main

import (
	"bufio"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 4096)
	for {
		reader := bufio.NewReader(conn)
		n, err := reader.Read(buf[:]) //读取数据

		if err != nil {
			println("读取数据失败，err：", err)
			break
		}
		s := string(buf[:n])
		println("收到的数据为:", s)
		conn.Write([]byte("返回成功"))
	}
}
func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:30000")

	if err != nil {
		println(err)
		return
	}
	println("正在接收信息")
	for {
		//等待客户端建立连接
		conn, err := listen.Accept() //监听

		if err != nil {
			println("conn err:", err)
			continue
		}

		go process(conn)
	}
}
