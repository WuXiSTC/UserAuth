package Service

import "../Dao"

func NewUser(ID, PASS string) (bool, error) {
	user := new(User)
	user.SetID(ID)
	user.SetPASS(PASS)
	ok, err := Dao.QueryUser(user)
	if err == nil && ok {
		return false, nil
	}
	return Dao.InsertUser(user) //不管什么错误还是成功不成功一并返回即可
}
