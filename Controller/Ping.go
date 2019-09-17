package Controller

import (
	"../Dao/Cache"
	"github.com/kataras/iris"
)

/*从机模式下用于从外部探测服务是否在线*/
func Ping(ctx iris.Context) {
	pong, err := Cache.Ping()
	if err == nil {
		logResponse(ctx.WriteString(pong))
	} else {
		logResponse(ctx.WriteString(""))
	}
}
