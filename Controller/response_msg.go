package Controller

import (
	"github.com/kataras/iris"
	"log"
)

func responseMsg(ctx iris.Context, data iris.Map) {
	state, err := ctx.JSON(data)
	if err != nil {
		log.Println(state, err)
	}
}
