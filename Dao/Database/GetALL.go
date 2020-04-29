package Database

import (
	"gitee.com/WuXiSTC/UserAuth/util"
)

/*获取数据库中指定量的键值对*/
func GetALL(n uint64) (map[string]string, error) {
	if Conf.UseCassandra {
		return cdbGetALL(n)
	}
	return rdbGetALL(n)
}

func rdbGetALL(n uint64) (map[string]string, error) {
	db, err := RDBConnect()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT ID,PASS FROM Users LIMIT ?", n)
	if err != nil {
		return nil, err
	}
	all := make(map[string]string)
	for rows.Next() {
		var ID, PASS string
		err = rows.Scan(&ID, &PASS)
		if err != nil {
			return all, err
		}
		all[ID] = PASS
	}
	return all, nil
}

func cdbGetALL(n uint64) (map[string]string, error) {
	db, err := CDBConnect()
	if err != nil {
		return nil, err
	}
	session, err := db.CreateSession()
	if err != nil {
		return nil, err
	}
	var ID, PASS string
	all := make(map[string]string)
	iter := session.Query("SELECT ID,PASS FROM Users LIMIT ?", n).Iter()
	for iter.Scan(&ID, &PASS) {
		all[ID] = PASS
	}
	util.LogE(iter.Close())
	session.Close()
	return all, nil
}
