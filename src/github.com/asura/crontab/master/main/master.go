package main

import (
	"flag"
	"fmt"
	"github.com/asura/crontab/master"
	"runtime"
	"time"
)

var (
	confFile string // 配置文件路径
)

// 解析命令行参数
func initArgs() {
	// master -config ./master.json
	flag.StringVar(&confFile, "config", "./master.json", "请指定 master.json 文件")

	flag.Parse()
}

// 初始化线程
func initMasterEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	var (
		err error
	)
	// 初始化命令行参数
	initArgs()
	// 初始化线程池
	initMasterEnv()

	// 加载配置
	if err = master.InitConfig(confFile); err != nil {
		goto ERR
	}

	// 日志管理器
	if err = master.InitLogMgr(); err != nil {
		goto ERR
	}

	// 任务管理器
	if err = master.InitJobMgr(); err != nil {
		goto ERR
	}

	// 启动Api HTTP服务
	if err = master.InitApiServer(); err != nil {
		goto ERR
	}

	// 正常退出
	for true {
		time.Sleep(1 * time.Second)
	}

	return

ERR:
	fmt.Println(err)
}
