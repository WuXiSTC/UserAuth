package Dao

import "./Database"
import "./Cache"

/*在数据库中更新密码*/
func UpdateUser(user, newUser User) (bool, error) {
	ID := user.GetID()
	_, _ = Cache.DelUser(ID) //不管三七二十一先删缓存
	PASS := user.GetPASS()
	newPASS := newUser.GetPASS()
	ok, err := Database.UpdateUser(ID, PASS, newPASS)
	errorHandler(err)
	if err == nil && ok { //数据库的验证通过了就写入缓存
		_, errR := Cache.SetUser(ID, newPASS)
		errorHandler(errR)
	}
	return ok, err
}
