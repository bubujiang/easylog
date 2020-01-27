package http

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	Ip string
	Port uint64
	Pid string
}

func (s *Server) _start() (*gin.Engine, error) {
	return gin.Default(), nil
}

func (s *Server) _deploy(router *gin.Engine) error {
	router.Static("/dist", "html/dist")
	router.Static("/plugins", "html/plugins")
	//router.SetFuncMap(template.FuncMap{
	//	"ShowModulesTags": cfun.ShowModulesTags,
	//})
	router.LoadHTMLGlob("html/tpl/*")

	return nil
}

func (s *Server) _route(router *gin.Engine) error {
	router.GET("/", index)
	router.POST("/",index)
	return nil
}

func (s *Server) _run(router *gin.Engine) error {
	return router.Run(":8080")
}

func (s *Server) Start() {
	router,_  := s._start()
	s._route(router)
	s._run(router)
}
