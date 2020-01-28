package main

import (
	cli "github.com/jawher/mow.cli"
	mconfig "log-server/config"
	mhttp "log-server/sever/http"
	"os"
)

type actionType string
const (
	startLog      actionType = "start-log"
	startHttp     actionType = "start-http"
	startAll      actionType = "start"
	reStartLog    actionType = "restart-log"
	reStartHttp   actionType = "restart-http"
	reStartAll    actionType = "restart"
	stopLog       actionType = "stop-log"
	stopHttp      actionType = "stop-http"
	stopAll       actionType = "stop"
)
func (t actionType) verify() bool {
	switch t {
	case startLog: return true
	case startHttp: return true
	case startAll: return true
	case reStartHttp: return true
	case reStartLog: return true
	case reStartAll: return true
	case stopLog: return true
	case stopHttp: return true
	case stopAll: return true
	}
	return false
}

type Arguments struct {
	configFile string
	action     actionType
}

//var (Cnf *mconfig.Config)
func init() {

}
func main() {

	var (
		app = cli.App("mserver", "an log server")
		args Arguments
	)

	//ss:=os.Args
	//fmt.Printf("%+v",ss)

	app.Spec = "C [-a]"

	app.StringArgPtr(&args.configFile, "C", "conf.ini", "path to config file")
	//app.StringArgPtr(&args.configFile, "c", "", "path to config file")
	app.StringOptPtr((*string)(&args.action), "a", string(startHttp), "action to do")

	// Specify the action to execute when the app is invoked correctly
	app.Action = func() {
		if !args.action.verify() {
			panic("action error")
		}

		mconfig.Init("conf.ini")
		//mconfig.Init(args.configFile)

		switch args.action {
		case startLog:
		case startHttp:
			//todo 判断是否已经开始,如果已经开始直接抛出异常.
			//开启http服务
			(&mhttp.Server{
				Ip:mconfig.Cnf.HttpServer.Ip,
				Port:mconfig.Cnf.HttpServer.Port,
				Pid:mconfig.Cnf.HttpServer.Pid,
			}).Start()
		case startAll:
			//todo 判断是否已经开始,开始
		case reStartHttp:
		case reStartLog:
		case reStartAll:
			//todo 平滑重启
		case stopLog:
		case stopHttp:
		case stopAll:
		}
		//return false

		//fmt.Printf("%+v", args)
	}
	// Invoke the app passing in os.Args
	app.Run(os.Args)
}