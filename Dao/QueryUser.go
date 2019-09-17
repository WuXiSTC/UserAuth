package Dao

import (
	"../util"
	"./Cache"
	"./Database"
)

//验证用户：验证用户名密码是否正确
//
//如果发现验证的用户不在缓存就会去数据库查找，如果查到了就写缓存
//
//返回值是(bool,error)分别表示验证是否通过和错误信息
func QueryUser(user User) (bool, error) {
	ID := user.GetID()
	PASS := user.GetPASS()

	ok, err := VerifyCachedUser(user) //先查缓存
	if err == nil && ok {
		return true, nil //缓存中有则返回
	}

	cachedPASS, err := Database.GetUserPASS(ID) //否则查数据库
	util.LogE(err)
	if err == nil && cachedPASS != "" { //数据库查到了就写入缓存
		_, errR := Cache.SetUser(ID, PASS)
		util.LogE(errR)
	}
	return cachedPASS == PASS, err //最后返回
}
