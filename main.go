//main.go 主文件

package main

import (
	"./util"
	"github.com/kataras/iris"
)

type Config struct {
	Listen  string `yaml:"Listen"`
	SlaveMode bool   `yaml:"SlaveMode"`
}

var Conf = Config{":8080", false}

//main 主函数
func main() {

	var app *iris.Application

	util.GetConf("Config/ServerConfig.yaml", &Conf)
	if Conf.SlaveMode { //从机模式
		app = SlaveApp("Config/SlaveConfig.yaml")
	} else { //主机模式
		app = MasterApp(
			"Config/DatabaseConfig.yaml",
			"Config/RedisConfig.yaml",
			"Config/DaemonsConfig.yaml",
		)
	}

	err := app.Run(iris.Addr(Conf.Listen))
	util.LogE(err)
}
