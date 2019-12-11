package server

import (
	MHttp "log-server/server/http"
	MLog "log-server/server/log"
	)

type Server struct {
	Log MLog.Server
	Http MHttp.Server
}
