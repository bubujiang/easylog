package config

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"os"
)

type Config struct {
	Db string `json:"db"`
	DSN string `json:"dsn"`
	Database string `json:"database"`
	DbTable string `json:"table"`
	DbPools int `json:"db_pools"`

	LogIp string `json:"log_ip"`
	LogPort uint32 `json:"log_port"`

	//LogAllowModule []string `json:"allow_module"`
	//LogAllowTags []string `json:"allow_tags"`
	LogAllow interface{} `json:"allow_search"`

	Net string `json:"net"`
}


func (config *Config)GetConfByFile() {
	var path string
	flag.StringVar(&path, "f", "config.cnf", "set configuration `file`")
	//
	file, _ := os.Open(path)
	read := bufio.NewReader(file)
	buffer := bytes.NewBuffer([]byte{})
	for {
		b, _, err := read.ReadLine()
		if err != nil {
			if err == io.EOF{
				break
			}
			//panic(err)
		}
		buffer.Write(b)
	}
	ret := make([]byte, buffer.Len())
	buffer.Read(ret)
	defer file.Close()
	err:=json.Unmarshal(ret,&config)
	if err!=nil {}
	//return ret
}

//var GConf *Config
var GConf = new(Config)

