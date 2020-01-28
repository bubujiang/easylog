package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"strconv"
	"time"
)

type Mongo struct {
	client *mongo.Database
	DSN string
	Database string
}

func (db *Mongo) Connect() error{
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(db.DSN))
	db.client = client.Database(db.Database)
	if !db.Validate() {
		//todo 返回错误
		//return errors.Wrap(err, "read failed")
	}
	return nil
}

func (db *Mongo) Close() error{
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return db.client.Client().Disconnect(ctx)
}

func (db *Mongo) Validate() bool{
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := db.client.Client().Ping(ctx, readpref.Primary())
	if err!=nil {
		return false
	}
	return true
}



func (db *Mongo) Find(resCond map[string]interface{}) []map[string]interface{} {

	ctx, _ := context.WithTimeout(context.Background(), 30000000*time.Second)
	collection := db.client.Collection(resCond["module"].(string))
	page,_ := strconv.ParseInt(resCond["page"].(string),10,64)
	num := resCond["num"].(int64)
	//if a != nil {}
	skipNum := (page - 1) * num

	cond := mkCondition(resCond)

	cur, err := collection.Find(ctx, cond, options.Find().SetLimit(num), options.Find().SetSkip(skipNum), options.Find().SetSort(bson.M{"time": -1}))
	if err != nil { log.Fatal(err) }
	defer cur.Close(ctx)
	var result []map[string]interface{}
	for cur.Next(ctx) {
		var row map[string]interface{}
		if err = cur.Decode(&row); err != nil {
			log.Fatal(err)
		}
		result = append(result, row)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func (db *Mongo) Insert() {

}

func (db *Mongo) Total(resCond map[string]interface{}) int64 {

	ctx, _ := context.WithTimeout(context.Background(), 30000000*time.Second)
	collection := db.client.Collection(resCond["module"].(string))

	cond := mkCondition(resCond)

	total,_ := collection.CountDocuments(ctx,cond)
	return total
}

func mkCondition(resCond map[string]interface{}) map[string]interface{} {
	f := make(map[string]interface{})
	f["tags"] = resCond["tags"]
	startTime, oks := resCond["start_time"]
	endTime, oke := resCond["end_time"]
	if oks && oke {
		timeStart := startTime.(int64)
		timeEnd := endTime.(int64)
		f["time"] = map[string]int64{"$gte": timeStart, "$lte": timeEnd}
	}else if oks {
		timeStart := startTime.(int64)
		f["time"] = map[string]int64{"$gte": timeStart}
	}else if oke{
		timeEnd := endTime.(int64)
		f["time"] = map[string]int64{"$lte": timeEnd}
	}

	return f
}
