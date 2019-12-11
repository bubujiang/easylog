package log

import (
	"fmt"
	"log-server/config"
	"net"
	"strconv"
)

type Server struct {
	IP string `json:"log_ip"`
	Port uint32 `json:"log_port"`
	Net string `json:"net"`
}

func (server *Server) _init() {
	server.IP = config.GConf.LogIp
	server.Port = config.GConf.LogPort
	server.Net = config.GConf.Net
}

func (server *Server) _start(handle func(net.Conn)) {
	tcpAddr,_ := net.ResolveTCPAddr(server.Net, server.IP+":"+strconv.FormatUint(uint64(server.Port), 10))
	tcpListener,_ := net.ListenTCP(server.Net,tcpAddr)
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
	server._init()
	server._start(_handle)

	fmt.Println("log start")
}
