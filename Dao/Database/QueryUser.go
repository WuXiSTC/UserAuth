package Database

/*查询用户是否存在*/
func QueryUser(ID, PASS string) (bool, error) {
	db, err := RDBConnect()
	if err != nil {
		return false, err
	}
	rows, err := db.Query("SELECT count(*) FROM Users WHERE ID=? AND PASS=? LIMIT 1", ID, PASS)
	if err != nil {
		return false, err
	}
	count := 0
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return false, err
		}
	}
	return count > 0, nil
}
