package Daemons

import (
	"gitee.com/WuXiSTC/UserAuth/Dao/Cache"
	"gitee.com/WuXiSTC/UserAuth/Dao/Interfaces"
	"gitee.com/WuXiSTC/UserAuth/util"
)

var DeleteQ = make(chan Interfaces.User, Conf.BufferSize)

//负责插入失败时在缓存中删除用户的守护进程
func DeleteDaemon() {
	for {
		user := <-DeleteQ
		_, err := Cache.DelUser(user.GetID())
		util.LogE(err)
	}
}
