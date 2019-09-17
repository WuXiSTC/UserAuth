package Database

/*获取用户密码*/
func GetUserPASS(ID string) (string, error) {
	db, err := RDBConnect()
	if err != nil {
		return "", err
	}
	rows, err := db.Query("SELECT PASS FROM Users WHERE ID=? LIMIT 1", ID)
	if err != nil {
		return "", err
	}
	var PASS = ""
	err = nil
	for rows.Next() {
		err = rows.Scan(&PASS)
	}
	return PASS, err
}
