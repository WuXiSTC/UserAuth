package Controller

import (
	"github.com/kataras/iris"
)

//在所有的数据操作接口之前调用的POST格式检查
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
