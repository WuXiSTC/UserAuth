package Daemons

import (
	"../../util"
	"../Database"
	"../Interfaces"
)

var InsertQ = make(chan Interfaces.User, Conf.BufferSize)

//负责插入用户的守护进程
func InsertDeamon() {
	for {
		user := <-InsertQ
		ok, err := Database.InsertUser(user.GetID(), user.GetPASS())
		if err != nil || !ok {
			util.LogE(err)
			DeleteQ <- user
		} else {
			util.Log("注册用户" + user.GetID() + "->" + user.GetPASS())
		}
	}
}
