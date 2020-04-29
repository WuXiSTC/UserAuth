package Dao

import (
	"gitee.com/WuXiSTC/UserAuth/Dao/Cache"
	"gitee.com/WuXiSTC/UserAuth/Dao/Database"
	"gitee.com/WuXiSTC/UserAuth/Dao/Interfaces"
	"gitee.com/WuXiSTC/UserAuth/util"
)

//查找用户是否存在：先查缓存再查数据库，数据库中有就放缓存
//
//有就返回true,nil；没有就返回false,[错误信息]
func UserExists(user Interfaces.User) (bool, error) {
	ID := user.GetID()
	PASS, err := Cache.GetUserPASS(ID)
	if err == nil && PASS != "" {
		return true, nil
	}
	PASS, err = Database.GetUserPASS(ID)
	if err == nil && PASS != "" {
		_, errR := Cache.SetUser(ID, PASS)
		util.LogE(errR)
		return true, nil
	}
	return false, err
}
