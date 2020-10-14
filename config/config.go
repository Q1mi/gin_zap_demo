package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config 整个项目的配置
type Config struct {
	Mode       string `json:"mode"`
	Port       int    `json:"port"`
	*LogConfig `json:"log"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
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
