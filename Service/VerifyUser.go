package Service

import "gitee.com/WuXiSTC/UserAuth/Dao"

//主机模式下的验证用户
func VerifyUser(ID, PASS string) (bool, error) {
	user := new(User)
	user.SetID(ID)
	user.SetPASS(PASS)
	return Dao.QueryUser(user) //不管什么错误还是成功不成功一并返回即可
}
