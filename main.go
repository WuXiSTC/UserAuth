package main

import (
	"./util"
	"github.com/kataras/iris"
)

type config struct {
	SlaveMode     bool   `yaml:"SlaveMode"`     //是否是从属模式
	MasterAddress string `yaml:"MasterAddress"` //如果是从属模式，则在此指定主服务的Redis地址
}

var Conf = config{false, ""}

func main() {

	util.GetConf("SlaveConfig.yaml", &Conf)

	var app *iris.Application

	if Conf.SlaveMode { //从属模式
		app = SlaveApp(Conf)
	} else { //主机模式
		app = MasterApp(Conf)
	}

	err := app.Run(iris.Addr(":8080"))
	util.LogE(err)
}
