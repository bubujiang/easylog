package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	IP string `json:"http_ip"`
	Port uint32 `json:"http_port"`
	Net string `json:"net"`
}

func (server *Server) init(json []byte) {
	fmt.Println(json)
}

func (server *Server) _start() {

}

func (server *Server) Start()  {
	router := gin.Default()
	router.Static("/dist", "html/dist")
	router.Static("/plugins", "html/plugins")
	router.LoadHTMLGlob("html/tpl/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/", index)
	router.POST("/",index)
	router.Run(":8080")
}