package main

import (
	"./Controller"
	"./Dao"
	"./util"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

func SlaveApp(conf Dao.SlaveConfig) *iris.Application {
	err := Dao.SlaveModeInit(conf)
	util.LogE(err)
	app := iris.New()
	app.Use(logger.New())
	app.Get("/PING", Controller.Ping) //PING
	app.Use(Controller.BeforeHandler)
	app.Post("/verify", Controller.VerifyCached) //验证用户名和密码
	return app
}
