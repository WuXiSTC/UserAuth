package Dao

import (
	"../util"
	"./Cache"
	"./Database"
	"./Interfaces"
)

//更新密码，先删缓存再更新数据库再写入缓存
//
//更新成功则返回true,nil；用户名密码错误则返回false,nil；出错则返回false,[错误信息]
func UpdateUser(newUser Interfaces.User) (bool, error) {
	ID := newUser.GetID()
	_, err := Cache.DelUser(ID) //不管三七二十一先删缓存
	util.LogE(err)

	newPASS := newUser.GetPASS()
	ok, err := Database.UpdateUser(ID, newPASS)
	if ok { //数据库修改完成了就写入缓存
		_, errR := Cache.SetUser(ID, newPASS)
		util.LogE(errR)
	}
	return ok, err
}
