package Daemons

import (
	"../../util"
	"../Database"
	"../Interfaces"
)

var InsertQ = make(chan Interfaces.User)

//负责插入用户的守护进程
func InsertDeamon() {
	user := <-InsertQ
	ok, err := Database.InsertUser(user.GetID(), user.GetPASS())
	if err != nil || !ok {
		util.LogE(err)
		DeleteQ <- user
	}
}
