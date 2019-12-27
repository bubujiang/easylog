package mmongo

import (
	"context"
	"log-server/config"

	//"fmt"
	//"strconv"
	//"sync/atomic"

	"github.com/jolestar/go-commons-pool/v2"
)

//type MyPoolObject struct {
//	s string
//}

type MyDbFactory struct {
	//v uint64
}

func (f *MyDbFactory) MakeObject(ctx context.Context) (*pool.PooledObject, error) {
	var mongo = new(MMongo)
	mongo.Init()
	mongo.Connect()
	return pool.NewPooledObject(mongo),nil
}

func (f *MyDbFactory) DestroyObject(ctx context.Context, object *pool.PooledObject) error {
	//object.Object.
	o := object.Object.(*MMongo)
	o.Close()
	// do destroy
	return nil
}

func (f *MyDbFactory) ValidateObject(ctx context.Context, object *pool.PooledObject) bool {
	// do validate
	o := object.Object.(*MMongo)
	return o.Validate()
}

func (f *MyDbFactory) ActivateObject(ctx context.Context, object *pool.PooledObject) error {
	// do activate
	return nil
}

func (f *MyDbFactory) PassivateObject(ctx context.Context, object *pool.PooledObject) error {
	// do passivate
	return nil
}

var MPools *pool.ObjectPool

func PoolsInit() {
	ctx := context.Background()
	MPools = pool.NewObjectPoolWithDefaultConfig(ctx, &MyDbFactory{})
	MPools.Config.MaxTotal = config.GConf.DbPools
}

//func Example_customFactory() {
//	ctx := context.Background()
//	p := pool.NewObjectPoolWithDefaultConfig(ctx, &MyDbFactory{})
//
//	obj1, err := p.BorrowObject(ctx)
//	if err != nil {
//		panic(err)
//	}
//
//	o := obj1.(*MMongo)
//	fmt.Println(o.s)
//
//	err = p.ReturnObject(ctx, obj1)
//	if err != nil {
//		panic(err)
//	}
//
//	// Output: 1
//}


//package mmongo
//
//import (
//	"log-server/config"
//	//"go.mongodb.org/mongo-driver/mongo"
//	"log-server/db"
//	"sync"
//)
//
//type Pools struct {
//	Clients []db.DB
//	Max uint32
//}
//
//func (pools *Pools)Init()  {
//	pools.Max = config.GConf.DbPools
//	//pools.Clients = make([]db.DB,pools.Max)
//}
//
//func (pools *Pools) GetClient() db.DB {
//	var m sync.Mutex
//	m.Lock()
//	defer m.Unlock()
//	//var mongo db.DB
//	AGAIN://todo 可能一直循环获取,需要加上超时报错退出逻辑
//	//从头遍历,找到可用的,直接返回
//	for _, v := range pools.Clients {
//		if v==nil {
//			continue
//		}
//		m,_ := v.(*MMongo)
//		if m.flag == AVAILABLE {
//			m.flag = USED
//			//mongo = v
//			return v
//		}
//	}
//	//
//	if pools.Max > uint32(len(pools.Clients)) {
//		//新建并放入
//		var mongo = new(MMongo)
//		mongo.Init()
//		mongo.Connect()
//		pools.PutClient(mongo)
//
//		return mongo
//	}
//	goto AGAIN
//	m.Unlock()
//	return nil
//}
//
//func (pools *Pools) PutClient(mongo db.DB) {
//	m,_ := mongo.(*MMongo)
//	//k := len(pools.Clients)
//	m.flag = AVAILABLE
//	//pools.Clients[k] = m
//	pools.Clients = append(pools.Clients, m)
//}