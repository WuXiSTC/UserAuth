package Dao

import (
	"gitee.com/WuXiSTC/UserAuth/Dao/Cache"
	"gitee.com/WuXiSTC/UserAuth/Dao/Daemons"
	"gitee.com/WuXiSTC/UserAuth/Dao/Interfaces"
	"gitee.com/WuXiSTC/UserAuth/util"
)

//更新密码，先删缓存再更新数据库再写入缓存
//
//更新成功则返回true,nil；失败则返回false,[错误信息]
func UpdateUser(newUser Interfaces.User) (bool, error) {
	ID := newUser.GetID()
	PASS := newUser.GetPASS()
	_, err := Cache.SetUser(ID, PASS) //直接写缓存
	util.LogE(err)
	if err != nil { //如果缓存写入失败则回滚
		_, errR := Cache.DelUser(ID)
		util.LogE(errR)
		return false, err
	}
	Daemons.UpdateQ <- newUser
	return true, nil
}
