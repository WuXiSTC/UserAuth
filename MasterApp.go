package main

import (
	"./Controller"
	"./Dao"
	"./util"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

func MasterApp() *iris.Application {
	util.LogE(Dao.CacheInit())

	app := iris.New()
	app.Use(logger.New())
	app.Use(Controller.BeforeHandler)
	app.Post("/verify", Controller.Verify)     //验证用户名和密码
	app.Post("/register", Controller.Register) //注册
	app.Post("/update", Controller.Update)     //修改用户名密码
	return app
}