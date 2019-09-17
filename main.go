package main

import (
	"./Dao"
	"./util"
	"github.com/kataras/iris"
)

var slaveConf = Dao.SlaveConfig{}

func main() {

	util.GetConf("SlaveConfig.yaml", &slaveConf)

	var app *iris.Application

	if slaveConf.SlaveMode { //从属模式
		app = SlaveApp(slaveConf)
	} else { //主机模式
		app = MasterApp()
	}

	err := app.Run(iris.Addr(":8080"))
	util.LogE(err)
}
