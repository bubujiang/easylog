package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
		go handleConnection(conn)
	}
	//net.TCPAddr
	//fmt.Printf("yzj is of type %T\n", yzj)
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
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
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://admin:123456qq@cluster0-xxhko.azure.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
		// handle error
	}
	//ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
	}

	collection := client.Database("testing").Collection("numbers")
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	id := res.InsertedID
	fmt.Println(id)
	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:<123456qq>@cluster0-xxhko.azure.mongodb.net/test?retryWrites=true&w=majority"))
	//fmt.Println(string(json))
}
