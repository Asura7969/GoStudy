package config

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

func main() {
	config, err := config.NewConfig("ini", "./logcollect.conf")
	if err != nil {
		fmt.Println("new config failed,err:", err)
		return
	}

	port, err := config.Int("server::port")
	if err != nil {
		fmt.Println("read server port failed,err:", err)
		return
	}

	fmt.Println("port", port)

	log_level := config.String("log::log_level")
	if len(log_level) == 0 {
		log_level = "info"
	}
	fmt.Println("log_level:", log_level)

	log_path := config.String("log::log_path")
	fmt.Println("log_path", log_path)

}
