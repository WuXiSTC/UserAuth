package Dao

import (
	"../util"
	"./Cache"
	"./Database"
)

//更新密码，先删缓存再更新数据库再写入缓存
//
//返回值是(bool,error)分别表示更新是否成功和错误信息
func UpdateUser(user, newUser User) (bool, error) {
	ID := user.GetID()
	_, _ = Cache.DelUser(ID) //不管三七二十一先删缓存
	PASS := user.GetPASS()
	newPASS := newUser.GetPASS()
	ok, err := Database.UpdateUser(ID, PASS, newPASS)
	util.LogE(err)
	if err == nil && ok { //数据库的验证通过了就写入缓存
		_, errR := Cache.SetUser(ID, newPASS)
		util.LogE(errR)
	}
	return ok, err
}
