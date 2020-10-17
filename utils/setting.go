package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string // 应用模式
	HttpPort string // 网络端口
	JwtKey   string // token验证
	Db         string // 数据库
	DbHost     string // 数据库主机
	DbPort     string // 数据库端口
	DbUser     string // 用户名
	DbPassWord string // 密码
	DbName     string // 数据库名称
)

func init() {
	// 使用ini包读取配置文件
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，检查路径：", err)
	}
	LoadServer(file)
	LoadData(file)
}

// 服务端实例化配置
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("jwtkey").MustString("jwtkey")
}

// 数据库端实例化配置
func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("debug")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("hua")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("root")
	DbName = file.Section("database").Key("DbName").MustString("123")
}
