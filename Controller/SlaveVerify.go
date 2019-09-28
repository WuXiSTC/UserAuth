package Controller

import (
	"../Service"
	"../util"
	"github.com/kataras/iris"
)

//只用缓存验证用户，仅用于Slave模式的/verify，封装Service.SlaveVerifyUser
func SlaveVerify(ctx iris.Context) {
	ID := ctx.PostValue("ID")
	PASS := ctx.PostValue("PASS")
	ok, err := Service.SlaveVerifydUser(ID, PASS)
	if err != nil {
		util.LogE(err)
		responseMsg(ctx, iris.Map{"ok": false, "message": "验证失败，请联系管理员"})
		return
	}
	if ok {
		responseMsg(ctx, iris.Map{"ok": true, "message": "验证通过"})
	} else {
		responseMsg(ctx, iris.Map{"ok": false, "message": "用户名或密码错误"})
	}
}
