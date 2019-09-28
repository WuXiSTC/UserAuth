package Controller

import (
	"../Service"
	"../util"
	"github.com/kataras/iris"
)

//修改密码，对应/update，封装Service.UpdateUser
//
//先Service.VerifyUser验证通过再更新
func Update(ctx iris.Context) {
	ID := ctx.PostValue("ID")
	PASS := ctx.PostValue("PASS")
	newPASS := ctx.PostValue("newPASS")
	if newPASS == "" {
		responseMsg(ctx, iris.Map{
			"ok":      false,
			"message": "请在newPASS中放上您的新密码"})
		return
	}
	ok, err := Service.VerifyUser(ID, PASS)
	if err != nil {
		util.LogE(err)
		responseMsg(ctx, iris.Map{
			"ok":      false,
			"message": "验证失败，请联系管理员"})
		return
	}
	if !ok {
		responseMsg(ctx, iris.Map{
			"ok":      false,
			"message": "用户名或密码错误"})
		return
	}

	ok, err = Service.UpdateUser(ID, PASS, newPASS)
	if err != nil || !ok {
		responseMsg(ctx, iris.Map{
			"ok":      false,
			"message": "更新失败，请联系管理员"})
		return
	}
	responseMsg(ctx, iris.Map{
		"ok":      true,
		"message": "更新成功"})
}
