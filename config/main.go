package config

import (
	"encoding/json"
	"gopkg.in/ini.v1"
)

type Config struct {
	HttpServer struct{
		Ip string
		Port uint64
		Pid string
	}

	LogServer struct{
		Ip string
		Port uint64
		Pid string
	}

	Log struct{
		ModulesTags map[string][]string
	}

	DB struct{
		DSN string
		Database string
		Max int
	}
}

var Cnf *Config

func Init(file string) {
	cfg, err := ini.Load(file)
	if err != nil {
		panic("file is not exist")
	}
	//cnf := &Config{}
	Cnf.HttpServer.Ip = cfg.Section("http-server").Key("ip").String()
	Cnf.HttpServer.Port,err = cfg.Section("http-server").Key("port").Uint64()
	if err != nil {
		panic("log server port error")
	}
	Cnf.HttpServer.Pid = cfg.Section("http-server").Key("pid").String()

	Cnf.LogServer.Ip = cfg.Section("log-server").Key("ip").String()
	Cnf.LogServer.Port,err = cfg.Section("log-server").Key("port").Uint64()
	if err != nil {
		panic("http server port error")
	}
	Cnf.LogServer.Pid = cfg.Section("log-server").Key("pid").String()

	err = json.Unmarshal(([]byte)(cfg.Section("log").Key("modules_tags").String()),Cnf.Log.ModulesTags)
	if err != nil {
		panic("modules tags error")
	}

	Cnf.DB.DSN = cfg.Section("db").Key("dsn").String()
	Cnf.DB.Database = cfg.Section("db").Key("database").String()
	Cnf.DB.Max,err = cfg.Section("db").Key("max").Int()
	if err != nil {
		panic("db max conn error")
	}
}



