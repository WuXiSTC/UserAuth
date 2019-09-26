//main.go 主文件

package main

import (
	"./Dao"
	"./util"
	"github.com/kataras/iris"
)

//main 主函数
func main() {

	var app *iris.Application

	Conf := Dao.SlaveConf
	util.GetConf("SlaveConfig.yaml", &Conf)
	if Conf.SlaveMode { //从机模式
		app = SlaveApp("SlaveConfig.yaml")
	} else { //主机模式
		app = MasterApp("DatabaseConfig.yaml", "RedisConfig.yaml")
	}

	err := app.Run(iris.Addr(":8080"))
	util.LogE(err)
}
