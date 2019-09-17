package Controller

import (
	"../util"
	"github.com/kataras/iris"
)

func responseMsg(ctx iris.Context, data iris.Map) {
	logResponse(ctx.JSON(data))
}

func logResponse(state int, err error) {
	if err != nil {
		util.LogE(err)
	}
}
