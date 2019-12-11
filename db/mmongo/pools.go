package mmongo

import (
	"log-server/config"
	//"go.mongodb.org/mongo-driver/mongo"
	"log-server/db"
	"sync"
)

type Pools struct {
	Clients []db.DB
	Max uint32
}

func (pools *Pools)Init()  {
	pools.Max = config.GConf.DbPools
	//pools.Clients = make([]db.DB,pools.Max)
}

func (pools *Pools) GetClient() db.DB {
	var m sync.Mutex
	m.Lock()
	defer m.Unlock()
	//var mongo db.DB
	AGAIN://todo 可能一直循环获取,需要加上超时报错退出逻辑
	//从头遍历,找到可用的,直接返回
	for _, v := range pools.Clients {
		if v==nil {
			continue
		}
		m,_ := v.(*MMongo)
		if m.flag == AVAILABLE {
			m.flag = USED
			//mongo = v
			return v
		}
	}
	//
	if pools.Max > uint32(len(pools.Clients)) {
		//新建并放入
		var mongo = new(MMongo)
		mongo.Init()
		mongo.Connect()
		pools.PutClient(mongo)

		return mongo
	}
	goto AGAIN
	m.Unlock()
	return nil
}

func (pools *Pools) PutClient(mongo db.DB) {
	m,_ := mongo.(*MMongo)
	//k := len(pools.Clients)
	m.flag = AVAILABLE
	//pools.Clients[k] = m
	pools.Clients = append(pools.Clients, m)
}