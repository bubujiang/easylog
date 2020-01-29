package log

import (
	"fmt"
	"net"
	"strconv"
)

type Server struct {
	Ip string
	Port uint64
	Pid string
	Net string
}

/*func (server *Server) _init() {
	server.Ip = config.Cnf.LogServer.Ip
	server.Port = config.Cnf.LogServer.Port
	server.Pid = config.Cnf.LogServer.Pid
}*/

func (server *Server) _start(handle func(net.Conn)) {
	tcpAddr,err := net.ResolveTCPAddr(server.Net, server.Ip+":"+strconv.FormatUint(uint64(server.Port), 10))
	if err != nil {
		panic("addr err")
	}
	tcpListener,err := net.ListenTCP(server.Net,tcpAddr)
	if err != nil {
		panic("listen err")
	}
	defer tcpListener.Close()
	for{
		conn,err := tcpListener.AcceptTCP()
		if err!=nil {
			fmt.Println(err)
			continue
		}
		go handle(conn)
	}
}

func (server *Server) Start()  {
	//server._init()
	server._start(Handle)

	fmt.Println("log start")
}
