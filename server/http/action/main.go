package action

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	mconfig "log-server/config"
	"strconv"

	//pools "github.com/jolestar/go-commons-pool/v2"
	mdb "log-server/db/mongodb"
	"net/http"
	"strings"
	"time"
)

func Index (c *gin.Context) {
	showData := gin.H{
		"moduleTags": mconfig.Cnf.Log.ModulesTags,
	}
	if strings.ToLower(c.Request.Method) == "post" {
		data := search(c)
		for k ,v := range data { showData[k] = v }
	} else {}

	c.HTML(http.StatusOK, "index.html", showData)
}


func search (c *gin.Context) gin.H {
	module := c.DefaultPostForm("module","")
	tag := c.DefaultPostForm("tag","")
	startTime := c.DefaultPostForm("start_time","")
	page := c.DefaultQuery("page","1")
	num,err := strconv.ParseInt(c.DefaultQuery("num","10"), 0, 64)
	if err != nil{
		num = 10
	}
	num = 1
	endTime := c.DefaultPostForm("end_time","")

	if module == "" || tag == "" {
		return gin.H{}
	}
	f := map[string]interface{}{"module":module,"tags":tag,"page":page,"num":num}
	if startTime!=""{
		stamp, _ := time.ParseInLocation("2006/01/02 15:04:05", startTime, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
		f["start_time"] = stamp.Unix()
	}
	if endTime!=""{
		stamp, _ := time.ParseInLocation("2006/01/02 15:04:05", endTime, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
		f["end_time"] = stamp.Unix()
	}
	//
	mdb.Init()
	p := mdb.P
	ctx := context.Background()

	obj, err := p.BorrowObject(ctx)
	if err != nil {
		panic(err)
	}

	db := obj.(*mdb.Mongo)
	rows := db.Find(f)
	//fmt.Println(o.s)

	err = p.ReturnObject(ctx, obj)
	if err != nil {
		panic(err)
	}




	for k,v := range rows{
		//time2 := v["time"].(int32)
		//time3 := time2.(int64)
		rows[k]["time"] = time.Unix(int64(v["time"].(int32)), 0).Format("2006/01/02 15:04:05")
		s,_ := json.Marshal(v["content"])
		rows[k]["content"] = string(s)
		//rows[k]["content"],_ = json.Marshal(v["content"])
	}
	return gin.H{
		"start_time":startTime,
		"end_time":endTime,
		"module":module,
		"tag":tag,
		"rows":rows,
		"page":page,
		"pages":db.Total(f)/num,
	}
	//fmt.Println(o.s)
	//strTime := c.DefaultQuery("time","0")
	//if c.
	//time,_ := strconv.ParseInt(c.DefaultQuery("time","0"),10,32)

	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
}