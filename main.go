package main

import (
	cli "github.com/jawher/mow.cli"
	mconfig "log-server/config"
	"log-server/market"
	mhttp "log-server/server/http"
	mlog "log-server/server/log"
	"os"
)



func main() {

	var (
		app = cli.App("mserver", "an log server")
		args market.Arguments
	)

	app.Spec = "C [-a]"

	app.StringArgPtr(&args.ConfigFile, "C", "conf.ini", "path to config file")
	//app.StringArgPtr(&args.configFile, "c", "", "path to config file")
	//app.StringOptPtr((*string)(&args.Action), "a", string(market.StartHttp), "action to do")
	app.StringOptPtr((*string)(&args.Action), "a", string(market.StartLog), "action to do")

	// Specify the action to execute when the app is invoked correctly
	app.Action = func() {
		if !args.Action.Verify() {
			panic("action error")
		}

		mconfig.Init("conf.ini")
		//mconfig.Init(args.configFile)

		switch args.Action {
		case market.StartLog:
			//todo 判断是否已经开始,如果已经开始直接抛出异常.
			//开启log服务
			(&mlog.Server{
				Ip:mconfig.Cnf.LogServer.Ip,
				Port:mconfig.Cnf.LogServer.Port,
				Pid:mconfig.Cnf.LogServer.Pid,
				Net:mconfig.Cnf.LogServer.Net,
			}).Start()
		case market.StartHttp:
			//todo 判断是否已经开始,如果已经开始直接抛出异常.
			//开启http服务
			(&mhttp.Server{
				Ip:mconfig.Cnf.HttpServer.Ip,
				Port:mconfig.Cnf.HttpServer.Port,
				Pid:mconfig.Cnf.HttpServer.Pid,
			}).Start()
		case market.StartAll:
			//todo 判断是否已经开始,开始
		case market.ReStartHttp:
		case market.ReStartLog:
		case market.ReStartAll:
			//todo 平滑重启
		case market.StopLog:
		case market.StopHttp:
		case market.StopAll:
		}
		//return false

		//fmt.Printf("%+v", args)
	}
	// Invoke the app passing in os.Args
	app.Run(os.Args)
}