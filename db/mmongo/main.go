package mmongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"log-server/config"
	"strconv"
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
	//Table string `json:"table"`
	client *mongo.Database
	conn *mongo.Client
	//flag bool
}

func (db *MMongo)Init()  {
	db.DSN = config.GConf.DSN
	db.Database = config.GConf.Database
	//db.Table = config.GConf.DbTable
	//db.flag = AVAILABLE
}

func (db *MMongo) Validate() bool{
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := db.conn.Ping(ctx, readpref.Primary())
	if err!=nil {
		return false
	}
	return true
}

func (db *MMongo) Connect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(db.DSN))
	db.conn = client

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	_ = db.Validate()

	//db.client = client.Database(db.Database).Collection(db.Table)
	db.client = client.Database(db.Database)
}

func (db *MMongo) Insert(b interface{}) bool{
	//ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//m := make(map[string]interface{})
	//j, _ := json.Marshal(b)
	//json.Unmarshal(j, &m)
	//_,err := db.client.InsertOne(ctx, m)
	//if err!=nil{}
	//id := res.InsertedID
	return true
}

func (db *MMongo) Find(b map[string]interface{}) []map[string]interface{}{

	ctx, _ := context.WithTimeout(context.Background(), 30000000*time.Second)
	collection := db.client.Collection(b["module"].(string))
	//delete(b, "module")
	page,_ := strconv.ParseInt(b["page"].(string),10,64)
	num,_ := strconv.ParseInt(b["num"].(string),10,64)
	skipNum := (page - 1) * num
	//delete(b)
	//skipNum := (b["page"] - 1) *  strconv.ParseInt(b["num"].(string))

	f := make(map[string]interface{})
	f["tags"] = b["tags"]
	startTime, oks := b["start_time"]
	endTime, oke := b["end_time"]
	if oks && oke {
		timeStart := startTime.(int64)
		timeEnd := endTime.(int64)
		f["time"] = map[string]int64{"$gte": timeStart, "$lte": timeEnd}
		//f["time"] := int64s2 := int64s
	}else if oks {
		timeStart := startTime.(int64)
		f["time"] = map[string]int64{"$gte": timeStart}
	}else if oke{
		timeEnd := endTime.(int64)
		f["time"] = map[string]int64{"$lte": timeEnd}
	}



	cur, err := collection.Find(ctx, f, options.Find().SetLimit(num), options.Find().SetSkip(skipNum), options.Find().SetSort(bson.M{"time": -1}))
	if err != nil { log.Fatal(err) }
	defer cur.Close(ctx)
	var result []map[string]interface{}
	for cur.Next(ctx) {
		var row map[string]interface{}
		if err = cur.Decode(&row); err != nil {
			log.Fatal(err)
		}
		result = append(result, row)

		//var result bson.M
		//err := cur.Decode(&result)
		//if err != nil { log.Fatal(err) }
		// do something with result....
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}


	//var a []map[string]interface{} //int array with length 3
	//a[0] = b
	return result
}

func (db *MMongo) Count(b map[string]interface{}) int64{

	ctx, _ := context.WithTimeout(context.Background(), 30000000*time.Second)
	collection := db.client.Collection(b["module"].(string))

	f := make(map[string]interface{})
	f["tags"] = b["tags"]
	startTime, oks := b["start_time"]
	endTime, oke := b["end_time"]
	if oks && oke {
		timeStart := startTime.(int64)
		timeEnd := endTime.(int64)
		f["time"] = map[string]int64{"$gte": timeStart, "$lte": timeEnd}
		//f["time"] := int64s2 := int64s
	}else if oks {
		timeStart := startTime.(int64)
		f["time"] = map[string]int64{"$gte": timeStart}
	}else if oke{
		timeEnd := endTime.(int64)
		f["time"] = map[string]int64{"$lte": timeEnd}
	}

	total,_ := collection.CountDocuments(ctx,f)
	return total
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

