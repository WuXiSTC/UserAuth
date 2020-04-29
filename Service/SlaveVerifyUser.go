package Service

import (
	"gitee.com/WuXiSTC/UserAuth/Dao"
)

//仅用于从机模式中验证用户
func SlaveVerifydUser(ID, PASS string) (bool, error) {
	user := new(User)
	user.SetID(ID)
	user.SetPASS(PASS)
	return Dao.VerifyCachedUser(user)
}
