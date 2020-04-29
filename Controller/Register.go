package Controller

import (
	"gitee.com/WuXiSTC/UserAuth/Service"
	"gitee.com/WuXiSTC/UserAuth/util"
	"github.com/kataras/iris"
)

//用户注册，对应/register，封装Service.NewUser函数
func Register(ctx iris.Context) {
	ID := ctx.PostValue("ID")
	PASS := ctx.PostValue("PASS")
	ok, err := Service.NewUser(ID, PASS)
	if err != nil {
		util.LogE(err)
		responseMsg(ctx, iris.Map{
			"ok":      false,
			"message": "注册失败，请联系管理员"})
		return
	}
	if !ok {
		responseMsg(ctx, iris.Map{
			"ok":      false,
			"message": "用户名已存在"})
		return
	}
	responseMsg(ctx, iris.Map{"ok": true, "message": "注册成功"})
}
