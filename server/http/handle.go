package http

import (
	"context"
	"encoding/json"
	"log-server/config"
	"log-server/db/mmongo"
	"strings"
	"time"

	//"fmt"
	"github.com/gin-gonic/gin"
	//"log-server/db/mmongo"
	"net/http"
)

func index (c *gin.Context) {
	//strTime := c.DefaultQuery("time","0")
	//if c.
	//time,_ := strconv.ParseInt(c.DefaultQuery("time","0"),10,32)
	showData := gin.H{
		"moduleTags": config.GConf.LogAllow,
	}
	if strings.ToLower(c.Request.Method) == "post" {
	//isPost := c.DefaultPostForm("search","")
	//if isPost=="search" {
		data := search(c)
		for k ,v := range data { showData[k] = v }
	} else {}
	//isPost = bool(isPost)
	//nick := c.DefaultPostForm("nick", "anonymous")

	c.HTML(http.StatusOK, "index.html", showData)
}


func search (c *gin.Context) gin.H {
	module := c.DefaultPostForm("module","")
	tag := c.DefaultPostForm("tag","")
	startTime := c.DefaultPostForm("start_time","")
	page := c.DefaultPostForm("page","1")
	num := c.DefaultPostForm("num","10")
	//endTime := c.DefaultPostForm("end_time","")

	//fmt.Printf(startTime)
	//fmt.Printf(endTime)
	//
	if module == "" || tag == "" {
		return gin.H{}
	}
	f := map[string]interface{}{"module":module,"tags":tag,"page":page,"num":num}
	if startTime!=""{
		stamp, _ := time.ParseInLocation("2006/01/02 15:04:05", startTime, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
		f["start_time"] = stamp.Unix()
	}
	//if endTime!=""{
	//	stamp, _ := time.ParseInLocation("2006/01/02 15:04:05", endTime, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	//	f["end_time"] = stamp.Unix()
	//}
	//
	ctx := context.Background()
	obj, err := mmongo.MPools.BorrowObject(ctx)
	if err != nil {
		panic(err)
	}
	db := obj.(*mmongo.MMongo)
	//return gin.H{}
	rows := db.Find(f)

	for k,v := range rows{
		//time2 := v["time"].(int32)
		//time3 := time2.(int64)
		rows[k]["time"] = time.Unix(int64(v["time"].(int32)), 0).Format("2006/01/02 15:04:05")
		s,_ := json.Marshal(v["content"])
		rows[k]["content"] = string(s)
		//rows[k]["content"],_ = json.Marshal(v["content"])
	}
	return gin.H{
		"module":module,
		"rows":rows,
		"pages":db.Count(f)/num
	}
	//fmt.Println(o.s)
	//strTime := c.DefaultQuery("time","0")
	//if c.
	//time,_ := strconv.ParseInt(c.DefaultQuery("time","0"),10,32)

//	c.JSON(200, gin.H{
//		"message": "pong",
//	})
}
