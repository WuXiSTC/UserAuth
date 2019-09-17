//main.go 主文件

package main

import (
	"./Dao"
	"./util"
	"github.com/kataras/iris"
)

var slaveConf = Dao.SlaveConfig{}

//main 主函数
func main() {

	util.GetConf("SlaveConfig.yaml", &slaveConf) //先读配置文件

	var app *iris.Application

	if slaveConf.SlaveMode { //从机模式
		app = SlaveApp(slaveConf)
	} else { //主机模式
		app = MasterApp()
	}

	err := app.Run(iris.Addr(":8080"))
	util.LogE(err)
}
