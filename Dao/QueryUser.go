package Dao

import (
	"../util"
	"./Cache"
	"./Database"
)

//验证用户：验证用户名密码是否正确，如果发现验证的用户不在缓存就会去数据库查找，如果查到了就写缓存
//
//如果验证通过则返回true,nil；用户名密码错误则返回false,nil；出错则返回false,[错误信息]
func QueryUser(user User) (bool, error) {
	ok, err := VerifyCachedUser(user) //先查缓存
	if err == nil {
		return ok, nil //缓存中有则返回
	}
	util.LogE(err)

	ID := user.GetID()
	PASS := user.GetPASS()
	cachedPASS, err := Database.GetUserPASS(ID) //否则查数据库
	if err != nil {
		return false, err
	}
	if cachedPASS != "" { //数据库查到了就写入缓存
		_, err := Cache.SetUser(ID, cachedPASS)
		util.LogE(err)
	}
	return cachedPASS != "" && cachedPASS == PASS, nil //最后返回
}
