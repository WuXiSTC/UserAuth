package Dao

import (
	"gitee.com/WuXiSTC/UserAuth/Dao/Cache"
	"gitee.com/WuXiSTC/UserAuth/Dao/Daemons"
	"gitee.com/WuXiSTC/UserAuth/Dao/Interfaces"
	"gitee.com/WuXiSTC/UserAuth/util"
)

//插入用户，先插数据库再插缓存；只负责插入不负责用户名冲突验证
//
//插入成功返回true,nil；失败返回false,[错误信息]
func InsertUser(user Interfaces.User) (bool, error) {
	ID := user.GetID()
	PASS := user.GetPASS()
	_, err := Cache.SetUser(ID, PASS) //直接写缓存
	util.LogE(err)
	if err != nil { //如果缓存写入失败则回滚
		_, errR := Cache.DelUser(ID)
		util.LogE(errR)
		return false, err
	}
	Daemons.InsertQ <- user //缓存写入成功则通过InsertDaemon写数据库
	return true, nil
}
