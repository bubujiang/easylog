package main

import (
	"log-server/config"
	MServer "log-server/server"
)



func main() {
	var server MServer.Server
	config.GConf.GetConfByFile()
	server.Log.Start()
	//server.Http.Start(conf)
}

//func main() {
//	var aa interface{}
//	ss,ff = aa.(int)
//
//	//var t int
//
//	var pools db.Pools
//
//	//x = t
//
//	y, ok = pools.(*mmongo.Pools)
//}


/*type DB interface {
	Connect() interface{}
	Insert()
	//Find(map[string]interface{}) []map[string]interface{}
}

type MMongo struct {
	DB
	DSN string
	Database string
	Table string
}

func (db *MMongo) Connect() interface{} {
	return db
}

func (db *MMongo) Insert(){
}*/

//func (db *MMongo) Find(b map[string]interface{}) []map[string]interface{}{
//	return []map['aa']interface{}
//}

/*func main() {
	var aa interface{}
	ss,ff = aa.(int)

	//var t int

	var x DB

	//x = t

	y, ok = x.(*MMongo)
}*/



/**
Insert([]byte) bool
	Find(map[string]interface{}) []map[string]interface{}
*/

