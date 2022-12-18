package main

import (
	"go-mall/conf"
	"go-mall/routes"
)

func main() {

	// 配置文件初始化
	conf.Init()

	// 路由
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)

}
