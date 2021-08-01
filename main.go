package main

import (
	"ziyue/core"
	"ziyue/global"
	"ziyue/initialize"
	"ziyue/route"
)

func main() {

	global.Z_VP = core.Viper("./config.yaml") // 初始化viper
	global.Z_DB = initialize.Gorm()           // gorm连接数据库
	global.Z_LOG = core.Zap()                 // 初始化zap日志库
	r := route.InitRouter()
	r.Run()
}
