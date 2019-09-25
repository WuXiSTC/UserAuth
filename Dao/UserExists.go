package Dao

import (
	"../util"
	"./Cache"
	"./Database"
)

//查找用户是否存在：先查缓存再查数据库，数据库中有就放缓存
//
//有就返回true,nil；没有就返回false,[错误信息]
func UserExists(user User) (bool, error) {
	ID := user.GetID()
	PASS, err := Cache.GetUserPASS(ID)
	if PASS != "" {
		return true, nil
	}
	util.LogE(err)
	PASS, err = Database.GetUserPASS(ID)
	if PASS != "" {
		_, errR := Cache.SetUser(ID, PASS)
		util.LogE(errR)
		return true, nil
	}
	return false, err
}
