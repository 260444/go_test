package main

import (
	"bufio"
	"net"
	"os"
	"strings"
)

func main() {
	//带上超时机制的Dial
	//net.Dial("tcp", "127.0.0.1:30000",2 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		println("拨号失败:", err)
	}

	defer conn.Close() //程序结束后关闭连接
	//控制台获取输入的字符
	reader := bufio.NewReader(os.Stdin)

	for {
		//读取直到第一次遇到delim字节
		println("请输入你要发送的信息")
		input, _ := reader.ReadString('\n') // 读取用户输入转换为字符串
		//\r 代表回车，也就是打印头归位，回到某一行的开头。
		//\n代表换行，就是走纸，下一行。
		//去除input中的\r\n
		inputInfo := strings.Trim(input, "\r\n")
		// 如果输入exit就退出
		//toupper转换为大写字符
		if strings.ToUpper(inputInfo) == "EXIT" {
			return
		}

		//发送数据
		n1, err := conn.Write([]byte(inputInfo))
		if err != nil {
			println(n1)
			return
		}
		//读取数据
		buf := [1024]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			println(err)
			return
		}
		println(string(buf[:n]))
	}
}
