package mmongo

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log-server/config"
	"time"
	//"log-server"
)

const(
	AVAILABLE = false
	USED = true
)

type MMongo struct {
	DSN string `json:"dsn"`
	Database string `json:"database"`
	Table string `json:"table"`
	client *mongo.Collection
	flag bool
}

func (db *MMongo)Init()  {
	db.DSN = config.GConf.DSN
	db.Database = config.GConf.Database
	db.Table = config.GConf.DbTable
	db.flag = AVAILABLE
}

func (db *MMongo) Connect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(db.DSN))

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	_ = client.Ping(ctx, readpref.Primary())

	db.client = client.Database(db.Database).Collection(db.Table)
}

func (db *MMongo) Insert(b interface{}) bool{
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	m := make(map[string]interface{})
	j, _ := json.Marshal(b)
	json.Unmarshal(j, &m)
	_,err := db.client.InsertOne(ctx, m)
	if err!=nil{}
	//id := res.InsertedID
	return true
}

func (db *MMongo) Find(b map[string]interface{}) []map[string]interface{}{
	var a []map[string]interface{} //int array with length 3
	a[0] = b
	return a
}

func (db *MMongo)Close()  {
	//db.client.
}

//func toDoc(v interface{}) (doc *bson.Document, err error) {
//	bson.EC.String("_id", "some-note-id")
//	data, err := bson.Marshal(v)
//	if err != nil {
//		return
//	}
//
//	err = bson.Unmarshal(data, &doc)
//	return
//}

