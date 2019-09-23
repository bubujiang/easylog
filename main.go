package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
)

const (
	ip   = "0.0.0.0"
	port = 9099
)

func main() {
	tcp_addr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:9099")
	//var tcpaddr net.TCPAddr
	//tcp_addr.IP = []byte(ip)
	//tcp_addr.Port = port

	tcp_listen, err := net.ListenTCP("tcp", tcp_addr)
	defer tcp_listen.Close()
	if err != nil {
		fmt.Println("listen error")
	}

	for {
		conn, err := tcp_listen.Accept()
		if err != nil {
			// handle error
		}
		defer conn.Close()
		go handleConnection(conn)
	}
	//net.TCPAddr
	//fmt.Printf("yzj is of type %T\n", yzj)
}

func handleConnection(conn net.Conn) {
	conn.Close()
	reader := bufio.NewReader(conn)
	var jsonBuf bytes.Buffer
	for {
		// 读取一行数据，交给后台处理
		line, isPrefix, err := reader.ReadLine()
		if len(line) > 0 {
			jsonBuf.Write(line)
			if !isPrefix {
				saveDB(jsonBuf.Bytes())
				jsonBuf.Reset()
			}
		}
		if err != nil {
			break
		}
	}
}

func saveDB(json []byte) {
	fmt.Println(string(json))
}
