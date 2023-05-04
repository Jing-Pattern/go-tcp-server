package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func HandelConnect(conn net.Conn) {
	defer conn.Close()
	// 获取连接的客户端网络地址
	addr := conn.RemoteAddr()
	fmt.Println(addr, "客户端连接地址，success")
	buf := make([]byte, 4096)
	for {
		// 用来保存读到数据的切片
		//读取客户端发送的数据
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("客户端关闭，断开连接")
			return
		}
		if err != nil {
			log.Fatal(err, "read 返回数据有问题")
			return
		}
		// 处理数据
		conn.Write([]byte(buf[:n]))
		fmt.Println("服务器回显", string(buf[:n]))
		//fmt.Println(conn.Write([]byte(buf[:n])))
	}
	//conn.Write()
	//conn.Read()
}
func main() {
	wg := sync.WaitGroup{}
	// 指定服务器通信协议
	listen, err := net.Listen("tcp", "127.0.0.1:8001")
	defer listen.Close()
	if err != nil {
		log.Fatal("listenTcp 错误信息", err)
		return
	}
	fmt.Println("服务器等待客户端建立链接")
	// 阻塞监听客户端链接请求,成功建立连接，返回通信用的conn

	for {
		wg.Add(1)
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("listen.Accept", err)
			return
		}

		fmt.Println("服务器与客户端建立连接")
		go HandelConnect(conn)
	}
	// 具体完成客户端和服务器的数据通信
	wg.Wait()
}
