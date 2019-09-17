package Controller

import (
	"../Dao/Cache"
	"github.com/kataras/iris"
)

func Ping(ctx iris.Context) {
	pong, err := Cache.Ping()
	if err == nil {
		logResponse(ctx.WriteString(pong))
	} else {
		logResponse(ctx.WriteString(""))
	}
}
