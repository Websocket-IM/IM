package main

import (
	"ginchat/route"
	"ginchat/utils"
)

func main() {
	// 开启路由
	r := route.Route()
	// 配置初始化
	utils.InitConfig()
	// zip日志库初始化
	utils.InitLogger()
	// mysql初始化
	utils.InitMysql()
	// redis初始化
	utils.InitRedis()
	r.Run(":6060")
}
