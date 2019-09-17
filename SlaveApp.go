package main

import (
	"./Controller"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

func SlaveApp(conf config) *iris.Application {
	app := iris.New()
	app.Use(logger.New())
	app.Get("/PING", Controller.Ping) //PING
	app.Use(Controller.BeforeHandler)
	app.Post("/verify", Controller.VerifyCached) //验证用户名和密码
	return app
}
