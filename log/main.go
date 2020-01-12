package log

import (
	"encoding/json"
	"log-server/db"
)

type Format struct {
	Module string `json:"module"`
	Tags []string `json:"tags"`
	Time uint32 `json:"time"`
	//Content struct{
	//	Url string `json:"url"`
	//	Query map[string]interface{} `json:"query"`
	//	Post map[string]interface{} `json:"post"`
	//	Header map[string]interface{} `json:"header"`
	//	Response interface{} `json:"resp"`
	//} `json:"content"`
	Content map[string]interface{} `json:"content"`
	client db.DB
}

func (log *Format)RClient() db.DB {
	return log.client
}

func (log *Format)Init(b []byte,db db.DB)  {
	err := json.Unmarshal(b,log)
	if err != nil {
		return
	}
	//从链接池获得数据库链接
	//var pools = new(mmongo.Pools)
	//pools.Init()
	log.client = db
}
func (log *Format) Add() bool {
	//m,_ := log.Client.(*mmongo.MMongo)
	log.client.Insert(log)
	return true
}

func (log *Format) Find()  {

}
