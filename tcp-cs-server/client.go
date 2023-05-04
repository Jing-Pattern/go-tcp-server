package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func HandleClient(conn net.Conn) {
	str := make([]byte, 4096)
	for {
		n, err := os.Stdin.Read(str)
		if err != nil {
			fmt.Println("os.Stdin.Read err", err)
		}
		if n == 0 {
			return
		}
		// 写给服务器
		conn.Write(str[:n])
	}

}
func BackServer(conn net.Conn) {
	buf := make([]byte, 4096)

	for {
		n, err := conn.Read(buf)
		if n == 0 {
			return
		}
		if err != nil {
			fmt.Println("conn write err", err)
			return
		}
		fmt.Println("客户端回显数据", string(buf[:n]))
	}
}
func main1() {
	conn, err := net.Dial("tcp", "127.0.0.1:8002")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	// 主动写数据给服务器
	go HandleClient(conn)
	go BackServer(conn)

}
