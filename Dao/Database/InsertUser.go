package Database

/*插入用户*/
func InsertUser(ID, PASS string) (bool, error) {
	db, err := RDBConnect()
	if err != nil {
		return false, err
	}
	_, err = db.Exec("INSERT into Users(ID,PASS)VALUES(?,?)", ID, PASS)
	return err == nil, err
}
