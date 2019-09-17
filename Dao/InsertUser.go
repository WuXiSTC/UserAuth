package Dao

import (
	"../util"
	"./Cache"
	"./Database"
)

/*插入用户*/
func InsertUser(user User) (bool, error) {
	ID := user.GetID()
	PASS := user.GetPASS()
	ok, err := Database.InsertUser(ID, PASS)
	util.LogE(err)
	if err == nil && ok { //如果在数据库中插入成功了就向缓存中写入用户信息
		_, errR := Cache.SetUser(ID, PASS)
		util.LogE(errR)
	}
	return ok, err
}
