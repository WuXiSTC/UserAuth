package Database

import (
	"../../util"
)

/*获取用户密码*/
func GetUserPASS(ID string) (string, error) {
	if Conf.UseCassandra {
		return cdbGetUserPASS(ID)
	}
	return rdbGetUserPASS(ID)
}

func rdbGetUserPASS(ID string) (string, error) {
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

func cdbGetUserPASS(ID string) (string, error) {
	db, err := CDBConnect()
	if err != nil {
		return "", err
	}
	session, err := db.CreateSession()
	if err != nil {
		return "", err
	}
	PASS := ""
	iter := session.Query("SELECT PASS FROM Users WHERE ID=? LIMIT 1", ID).Iter()
	for iter.Scan(&PASS) {
	}
	util.LogE(iter.Close())
	session.Close()
	return PASS, nil
}
