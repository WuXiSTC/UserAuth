package Dao

import "./Database"
import "./Cache"

/*查询用户是否存在*/
func QueryUser(user User) (bool, error) {
	ID := user.GetID()
	PASS := user.GetPASS()
	cachedPASS, err := Cache.GetUserPASS(ID) //先查缓存
	errorHandler(err)
	if err == nil {
		return cachedPASS == PASS, nil //缓存中有则返回
	}
	cachedPASS, err = Database.GetUserPASS(ID) //否则查数据库
	errorHandler(err)
	if err == nil && cachedPASS != "" { //数据库查到了就写入缓存
		_, errR := Cache.SetUser(ID, PASS)
		errorHandler(errR)
	}
	return cachedPASS == PASS, err //最后返回
}
