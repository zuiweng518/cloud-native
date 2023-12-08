package utils

import (
	"github.com/spf13/viper"
)

func GetConfigInfo() (string, string, string, string, string) {
	v := viper.New()
	v.AddConfigPath(".") // 添加配置文件搜索路径，点号为当前目录
	v.AddConfigPath("conf")
	v.SetConfigType("yaml")
	v.SetConfigName("config.yaml")
	v.ReadInConfig()
	host := v.Sub("db").GetString("host")
	port := v.Sub("db").GetString("port")
	dbname := v.Sub("db").GetString("dbname")
	username := v.Sub("db").GetString("username")
	password := v.Sub("db").GetString("password")
	return host, port, dbname, username, password
}
