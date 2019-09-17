package Service

import "../Dao"

//修改密码：直接搞Dao.UpdateUser，如果更新不成功会自己报错的
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
