package Daemons

import (
	"../../util"
	"../Database"
	"../Interfaces"
)

var UpdateQ = make(chan Interfaces.User, Conf.BufferSize)

//负责更新用户密码的守护进程
func UpdateDaemon() {
	for {
		user := <-UpdateQ
		ok, err := Database.UpdateUser(user.GetID(), user.GetPASS())
		if err != nil || !ok {
			util.LogE(err)
			DeleteQ <- user
		} else {
			util.Log("修改密码" + user.GetID() + "->" + user.GetPASS())
		}
	}
}
