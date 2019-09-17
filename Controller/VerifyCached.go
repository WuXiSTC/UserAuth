package Controller

import (
	"../Service"
	"github.com/kataras/iris"
)

/*只用缓存验证用户，仅用于Slave模式*/
func VerifyCached(ctx iris.Context) {
	ID := ctx.PostValue("ID")
	PASS := ctx.PostValue("PASS")
	ok, err := Service.VerifyCachedUser(ID, PASS)
	if err != nil {
		responseMsg(ctx, iris.Map{"exists": false, "ok": false, "message": "缓存中没有这个用户"})
		return
	}
	if ok {
		responseMsg(ctx, iris.Map{"exists": true, "ok": true, "message": "验证通过"})
	} else {
		responseMsg(ctx, iris.Map{"exists": true, "ok": false, "message": "用户名或密码错误"})
	}
}
