package Controller

import (
	"gitee.com/WuXiSTC/UserAuth/Dao/Cache"
	"github.com/kataras/iris"
)

//从机模式下用于从外部探测服务是否在线，封装Cache.Ping函数
//
//正常情况返回"PONG"有错就返回""
func Ping(ctx iris.Context) {
	pong, err := Cache.Ping()
	if err == nil {
		logResponse(ctx.WriteString(pong))
	} else {
		logResponse(ctx.WriteString(""))
	}
}
