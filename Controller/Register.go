package Controller

import (
	"../Service"
	"github.com/kataras/iris"
)

/*用户注册API*/
func Register(ctx iris.Context) {
	ID := ctx.PostValue("ID")
	PASS := ctx.PostValue("PASS")
	ok, err := Service.NewUser(ID, PASS)
	if err != nil {
		responseMsg(ctx, iris.Map{
			"ok":      false,
			"message": "注册失败，请联系管理员"})
		return
	}
	if !ok {
		responseMsg(ctx, iris.Map{"ok": false, "message": "用户名已存在"})
		return
	}
	responseMsg(ctx, iris.Map{"ok": true, "message": "注册成功"})
}
