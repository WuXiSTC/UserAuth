package main

import (
	"gitee.com/WuXiSTC/UserAuth/Controller"
	"gitee.com/WuXiSTC/UserAuth/Dao"
	"gitee.com/WuXiSTC/UserAuth/util"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

//此函数用于在主函数中创建从机模式的iris.Application
func SlaveApp(ConfPath string) *iris.Application {
	Dao.SlaveConfigure(ConfPath)
	err := Dao.SlaveModeInit()
	util.LogE(err)
	app := iris.New()
	app.Use(logger.New())
	app.Get("/PING", Controller.Ping) //PING
	app.Use(Controller.BeforeHandler)
	app.Post("/verify", Controller.SlaveVerify) //验证用户名和密码
	return app
}
