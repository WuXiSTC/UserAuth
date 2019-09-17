package Dao

import (
	"../util"
	"./Cache"
)

/*
仅在缓存中验证用户名和密码
验证成功返回(true,nil)
用户名密码错误返回(false,nil)
其他情况返回(false,[错误信息])
*/
func VerifyCachedUser(user User) (bool, error) {
	cachedPASS, err := Cache.GetUserPASS(user.GetID())
	util.LogE(err)
	if err != nil {
		return false, err
	}
	return cachedPASS == user.GetPASS(), err
}
