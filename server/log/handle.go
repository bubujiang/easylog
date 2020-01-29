package log

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"io"
	mdb "log-server/db/mongodb"
	//"log-server/log"
	"log-server/market"
	"net"
)

func Handle(conn net.Conn) {
	//获得数据
	msg := getMsg(conn)
	//添加数据
	addMsg(msg)

	//msg := _getMsg(conn)
	/*MMongo.Insert(msg)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	collection := client.Database("testing").Collection("numbers")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	id := res.InsertedID
	fmt.Print("log handle")*/
}

func getMsg(conn net.Conn) []byte {
	reader := bufio.NewReader(conn)
	buffer := bytes.NewBuffer([]byte{})
	//接收并返回消息
	for {

		b, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF{
				break
			}
			//panic(err)
		}
		buffer.Write(b)
	}

	msg := make([]byte, buffer.Len())
	buffer.Read(msg)

	return msg
}

func addMsg(msg []byte){
	//初始化一个日志对象
	logData := &market.LogFormat{}
	err := json.Unmarshal(msg,logData)
	if err != nil {
		return
	}

	mdb.Init()
	p := mdb.P
	ctx := context.Background()

	obj, err := p.BorrowObject(ctx)
	if err != nil {
		panic(err)
	}

	db := obj.(*mdb.Mongo)
	db.Insert(logData)
	//fmt.Println(o.s)

	err = p.ReturnObject(ctx, obj)
	if err != nil {
		panic(err)
	}

	//从链接池获得数据库链接
	//var pools = new(mmongo.Pools)
	//pools.Init()
	//logData.Init(msg,pools.GetClient())
	//logData.Client = pools.GetClient()
	//调用日志对象的添加日志功能
	//ogData.Add()
	//数据库链接放回链接池
	//pools.PutClient(logData.RClient())
}
