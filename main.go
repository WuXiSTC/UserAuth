package main

import (
	"./Controller"
	"./Dao"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"log"
)

func main() {
	err := Dao.CacheInit()
	if err != nil {
		log.Println(err)
	}

	app := iris.New()
	app.Use(logger.New())
	app.Use(Controller.BeforeHandler)
	app.Post("/verify", Controller.Verify)     //验证用户名和密码
	app.Post("/register", Controller.Register) //注册
	app.Post("/update", Controller.Update)     //修改用户名密码

	err = app.Run(iris.Addr(":8080"))
	if err != nil {
		log.Println(err)
	}
}
