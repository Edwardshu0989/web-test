package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)

var Env config
var PrivateKey []byte

type config struct {
	DSN string
}

func init() {
	v := viper.New()
	v.SetConfigType("yaml")
	v.AddConfigPath("./conf")
	v.SetConfigName("config")

	if err := v.ReadInConfig(); err != nil {
		fmt.Errorf("读取配置文件信息错误 %s", err.Error())
	}

	//也可以直接反序列化为Struct

	if err := v.Unmarshal(&Env); err != nil {
		fmt.Printf("err:%s", err)
	}
	v.SetDefault("DSN", "root:ai132818@tcp(localhost:3306)/web-test?charset=utf8mb4&parseTime=True&loc=Local")

	var file *os.File
	file, err := os.Open("./config/server_private.pem")
	if err != nil {
		fmt.Errorf("私钥加载错误 %s", err.Error())
	}
	defer file.Close()
	PrivateKey, err = ioutil.ReadAll(file)
	if err != nil {
		fmt.Errorf("读取私钥文件信息出错 %s", err.Error())
	}

}
