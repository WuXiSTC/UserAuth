package Database

/*在数据库中更新密码*/
func UpdateUser(ID, PASS, newPASS string) (bool, error) {
	db, err := RDBConnect()
	if err != nil {
		return false, err
	}
	_, err = db.Exec(
		"UPDATE Users SET PASS=? WHERE ID=? AND PASS=?", newPASS, ID, PASS)
	return err == nil, err
}
