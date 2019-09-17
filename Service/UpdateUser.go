package Service

import "../Dao"

func UpdateUser(ID, PASS, newPASS string) (bool, error) {
	user := new(User)
	user.SetID(ID)
	user.SetPASS(PASS)
	newUser := new(User)
	newUser.SetID(ID)
	newUser.SetPASS(newPASS)
	ok, err := Dao.UpdateUser(user, newUser)
	if err != nil {
		return false, err //出错则返回错误信息
	}
	if !ok {
		return false, nil //Update不成功则返回false
	}
	return true, nil //Update成功则返回true
}
