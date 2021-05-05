package main

import (
	"ziyue/core"
	"ziyue/global"
	"ziyue/initialize"
	"ziyue/route"
)

func main() {

	global.ZVP = core.Viper("./config.yaml") // 初始化viper
	global.ZDB = initialize.Gorm()
	r := route.InitRouter()
	r.Run()
}
