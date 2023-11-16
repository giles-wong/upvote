package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config 整个项目的配置
type Config struct {
	Mode         string `json:"mode"`
	Port         int    `json:"port"`
	*LogConfig   `json:"logger"`
	*MySqlConfig `json:"mysql"`
}

// MySqlConfig mysql数据库配置
type MySqlConfig struct {
	User    string `json:"user"`
	Pass    string `json:"pass"`
	Host    string `json:"host"`
	Port    string `json:"port"`
	DbName  string `json:"dbName"`
	Charset string `json:"charset"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string `json:"level"`      //日志级别
	Path       string `json:"path"`       //日志目录
	MaxSize    int    `json:"maxsize"`    //单个文件最大 M
	MaxAge     int    `json:"maxAge"`     //备份文件存储多少天
	MaxBackups int    `json:"maxBackups"` //最多存储多少个备份文件
	Compress   bool   `json:"compress"`   //是否压缩
}

// Conf 全局配置变量
var Conf = new(Config)

// Init 初始化配置；从指定文件加载配置文件
func Init(filePath string) error {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, Conf)
}
