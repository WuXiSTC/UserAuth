package Dao

import (
	"../util"
	"./Cache"
	"./Database"
)

/*查询用户名密码是否正确*/
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
