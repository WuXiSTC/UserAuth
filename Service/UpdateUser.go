package Service

import "../Dao"

//修改密码：先用Dao.QueryUser进行验证再进行Dao.UpdateUser
//
//更新成功则返回true,nil；用户名密码错误则返回false,nil；出错则返回false,[错误信息]
func UpdateUser(ID, PASS, newPASS string) (bool, error) {
	user := new(User)
	user.SetID(ID)
	user.SetPASS(PASS)
	ok, err := Dao.QueryUser(user) //先验证用户名密码
	if !ok {
		return false, err //用户名密码不对直接就返回
	}
	newUser := new(User)
	newUser.SetID(ID)
	newUser.SetPASS(newPASS)
	return Dao.UpdateUser(newUser) //用户名密码对了才更新
}
