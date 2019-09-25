package Database

//在数据库中更新密码：仅负责更新密码不进行验证操作
//
//更新成功则返回true,nil；失败则返回false,[错误信息]
func UpdateUser(ID, newPASS string) (bool, error) {
	if Conf.UseCassandra {
		return cdbUpdateUser(ID, newPASS)
	}
	return rdbUpdateUser(ID, newPASS)
}

func rdbUpdateUser(ID, newPASS string) (bool, error) {
	db, err := RDBConnect()
	if err != nil {
		return false, err
	}
	_, err = db.Exec(
		"UPDATE Users SET PASS=? WHERE ID=?", newPASS, ID)
	return err == nil, err
}

func cdbUpdateUser(ID, newPASS string) (bool, error) {
	db, err := CDBConnect()
	if err != nil {
		return false, err
	}
	session, err := db.CreateSession()
	if err != nil {
		return false, err
	}
	err = session.Query(
		"UPDATE Users SET PASS=? WHERE ID=?", newPASS, ID).Exec()
	session.Close()
	return err == nil, err
}
