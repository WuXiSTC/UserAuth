package Controller

import (
	"../Service"
	"github.com/kataras/iris"
)

/*
仅用于Slave模式
一个特殊的缓存操作：只用缓存验证用户
在这样的调用只能在缓存结构没有隐藏时才可用，因此Service层无意义，直接从Controller层调用Dao层
*/
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
