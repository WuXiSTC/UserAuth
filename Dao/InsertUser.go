package Dao

import (
	"../util"
	"./Cache"
	"./Database"
)

//插入用户，先插数据库再插缓存；只负责插入不负责用户名冲突验证
//
//插入成功返回true,nil；失败返回false,[错误信息]
func InsertUser(user User) (bool, error) {
	ID := user.GetID()
	PASS := user.GetPASS()
	ok, err := Database.InsertUser(ID, PASS)
	if err == nil && ok { //如果在数据库中插入成功了就向缓存中写入用户信息
		_, errR := Cache.SetUser(ID, PASS)
		util.LogE(errR)
	}
	util.LogE(err)
	return ok, err
}
