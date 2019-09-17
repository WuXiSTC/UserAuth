package Controller

import (
	"github.com/kataras/iris"
)

func BeforeHandler(ctx iris.Context) {
	ID := ctx.PostValue("ID")
	PASS := ctx.PostValue("PASS")
	if ID == "" || PASS == "" {
		responseMsg(ctx, iris.Map{
			"ok":      false,
			"message": "请使用POST方法：在ID中放您的用户名、PASS中放您的密码"})
		return
	}
	ctx.Next()
}
