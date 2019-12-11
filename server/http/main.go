package http

import "fmt"

type Server struct {

}

func (server *Server) init(json []byte) {
	fmt.Println(json)
}

func (server *Server) _start() {

}

func (server *Server) Start(json []byte)  {
	server.init(json)
	server._start()

	fmt.Println("http start")
	//return
}