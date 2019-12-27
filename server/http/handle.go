package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func index (c *gin.Context) {
	//strTime := c.DefaultQuery("time","0")
	//if c.
	//time,_ := strconv.ParseInt(c.DefaultQuery("time","0"),10,32)
	isPost := c.DefaultPostForm("search","")
	if isPost=="search" {

	} else {}
	//isPost = bool(isPost)
	nick := c.DefaultPostForm("nick", "anonymous")

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}


func search (c *gin.Context) {
	module := c.DefaultPostForm("module","")
	tag := c.DefaultPostForm("tag","")
	start_time := c.DefaultPostForm("start_time","")
	end_time := c.DefaultPostForm("end_time","")
	//

	//strTime := c.DefaultQuery("time","0")
	//if c.
	//time,_ := strconv.ParseInt(c.DefaultQuery("time","0"),10,32)

//	c.JSON(200, gin.H{
//		"message": "pong",
//	})
}
