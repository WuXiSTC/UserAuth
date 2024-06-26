package main

import (
	"gitee.com/WuXiSTC/UserAuth/Controller"
	"gitee.com/WuXiSTC/UserAuth/Dao"
	"gitee.com/WuXiSTC/UserAuth/Dao/Cache"
	"gitee.com/WuXiSTC/UserAuth/Dao/Daemons"
	"gitee.com/WuXiSTC/UserAuth/Dao/Database"
	"gitee.com/WuXiSTC/UserAuth/util"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

//此函数用于在主函数中创建主机模式的iris.Application
func MasterApp(DbConfPath, RdsConfPath, DmsConfPath string) *iris.Application {
	Database.ConfigureDatabase(DbConfPath)
	Cache.ConfigureRedis(RdsConfPath)
	if Cache.Conf.InitAtStart {
		util.LogE(Dao.CacheInit())
	}
	Daemons.StartDaemons(DmsConfPath)
	app := iris.New()
	app.Use(logger.New())
	app.Use(Controller.BeforeHandler)
	app.Post("/verify", Controller.Verify)     //验证用户名和密码
	app.Post("/register", Controller.Register) //注册
	app.Post("/update", Controller.Update)     //修改用户名密码
	return app
}
