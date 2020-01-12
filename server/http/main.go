package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"log-server/html/cfun"
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
	router := gin.Default()
	router.Static("/dist", "html/dist")
	router.Static("/plugins", "html/plugins")
	router.SetFuncMap(template.FuncMap{
		"ShowModulesTags": cfun.ShowModulesTags,
	})
	router.LoadHTMLGlob("html/tpl/*")
	router.GET("/", index)
	router.POST("/",index)
	router.Run(":8080")
}

func (server *Server) Start()  {
	server._start()
}