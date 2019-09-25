package Database

//插入用户
//
//插入成功返回true,nil；失败返回false,[错误信息]
func InsertUser(ID, PASS string) (bool, error) {
	if Conf.UseCassandra {
		return cdbInsertUser(ID, PASS)
	}
	return rdbInsertUser(ID, PASS)
}

func rdbInsertUser(ID, PASS string) (bool, error) {
	db, err := RDBConnect()
	if err != nil {
		return false, err
	}
	_, err = db.Exec("INSERT into Users(ID,PASS)VALUES(?,?)", ID, PASS)
	return err == nil, err
}

func cdbInsertUser(ID, PASS string) (bool, error) {
	db, err := CDBConnect()
	if err != nil {
		return false, err
	}
	session, err := db.CreateSession()
	if err != nil {
		return false, err
	}
	err = session.Query("INSERT into Users(ID,PASS)VALUES(?,?)", ID, PASS).Exec()
	session.Close()
	return err == nil, err
}
