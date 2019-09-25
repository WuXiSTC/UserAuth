package Service

import (
	"../Dao"
)

//新增用户：先Dao.UserExists查用户存不存在再Dao.InsertUser
//
//如果成功返回true,nil；用户名冲突返回false,nil；失败返回false,[错误信息]
func NewUser(ID, PASS string) (bool, error) {
	user := new(User)
	user.SetID(ID)
	user.SetPASS(PASS)
	exists, err := Dao.UserExists(user) //先查用户是否存在
	if exists || err != nil {           //用户已存在或者报错就返回失败
		return false, err
	}
	return Dao.InsertUser(user) //否则就插入数据库，不管什么错误还是成功不成功一并返回即可
}
