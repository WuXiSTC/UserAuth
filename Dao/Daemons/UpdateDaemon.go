package Daemons

import (
	"../../util"
	"../Database"
	"../Interfaces"
)

var UpdateQ = make(chan Interfaces.User)

//负责更新用户密码的守护进程
func UpdateDaemon() {
	user := <-UpdateQ
	ok, err := Database.UpdateUser(user.GetID(), user.GetPASS())
	if err != nil || !ok {
		util.LogE(err)
		DeleteQ <- user
	}
}
