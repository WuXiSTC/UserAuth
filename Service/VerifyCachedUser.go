package Service

import (
	"../Dao"
)

//仅用于从机模式中，用缓存验证用户
func VerifyCachedUser(ID, PASS string) (bool, error) {
	user := new(User)
	user.SetID(ID)
	user.SetPASS(PASS)
	return Dao.VerifyCachedUser(user)
}
