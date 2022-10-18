package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type LogConfig struct {
	LogDir   string `json:"log_dir"`
	LogLevel string `json:"log_level"`
}

// 读取配置文件
func LoadLogConfig() *LogConfig {

	log_config := LogConfig{}

	open, err := os.Open("log/log_config.json")
	if err != nil {
		panic(err)
	}

	defer open.Close()

	//用流读取文件内容
	all, err := io.ReadAll(open)
	if err != nil {
		panic(err)
	}
	//json.Unmarshal(all, &log_config)
	//Unmarshal将数据读到结构体，第一的参数json字符串
	err = json.Unmarshal(all, &log_config)
	if err != nil {
		panic(err)
	}
	return &log_config
}

// 初始化记录器一个实例
var Logrus = logrus.New()

func init() {
	//读取日志的配置文件
	log_conf := LoadLogConfig()

	file, err := os.OpenFile(log_conf.LogDir, os.O_APPEND|os.O_CREATE|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	//将上面打开的fle设置为 日志输出文件
	Logrus.Out = file

	//设置日志级别
	//定义一个map 存储日志级别
	l_l_m := map[string]logrus.Level{
		"trace": logrus.TraceLevel,
		"info":  logrus.InfoLevel,
	}
	Logrus.SetLevel(l_l_m[log_conf.LogLevel])

	//日志格式化
	logrus.SetFormatter(&logrus.TextFormatter{})

}

func main() {
	engine := gin.Default()
	Logrus.Info("sssss")
	engine.Run(":8080")
}
