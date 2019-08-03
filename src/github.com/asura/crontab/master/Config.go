package master

import (
	"encoding/json"
	"io/ioutil"
)

// 程序配置
type Config struct {
	ApiPort               int      `json:"apiPor"`
	ApiReadTimeOut        int      `json:"apiReadTimeOut"`
	ApiWriteTimeOut       int      `json:"apiWriteTimeOut"`
	EtcdEndPointPoints    []string `json:"etcdEndPoints"`
	EtcdDailTimeout       int      `json:"etcdDailTimeout"`
	WebRoot               string   `json:"webroot"`
	MongodbUri            string   `json:"mongodbUri"`
	MongodbConnectTimeout int      `json:"mongodbConnectTimeout"`
}

var (
	G_Config *Config
)

func InitConfig(fileName string) (err error) {

	var (
		content []byte
		config  Config
	)
	// 把配置文件读进来
	if content, err = ioutil.ReadFile(fileName); err != nil {
		return
	}

	// json 反序列化
	if err = json.Unmarshal(content, &config); err != nil {
		return
	}
	// 赋值单例
	G_Config = &config

	//fmt.Println(config)

	return

}
