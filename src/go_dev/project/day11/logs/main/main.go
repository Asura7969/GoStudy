package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func main() {
	config := make(map[string]interface{})
	config["fileName"] = "./log/logcollect.log"
	config["level"] = logs.LevelDebug

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed,err:", err)
		return
	}

	logs.SetLogger(logs.AdapterAliLS, string(configStr))

	logs.Debug("this is a Debug,my name is %s", "stud01")
	logs.Trace("this is a Trace,my name is %s", "stud02")
	logs.Warn("this is a Warn,my name is %s", "stud03")

}
