package Dao

import (
	"./Cache"
	"./Database"
)

func CacheInit() error {
	_, _ = Cache.Clear()
	db, err := Database.RDBConnect()
	if err != nil {
		return err
	}
	rows, err := db.Query("SELECT ID,PASS FROM Users LIMIT ?", Cache.Conf.InitialCache)
	if err != nil {
		return err
	}
	for rows.Next() {
		var ID, PASS string
		err = rows.Scan(&ID, &PASS)
		if err != nil {
			return err
		}
		_, _ = Cache.SetUser(ID, PASS)
	}
	return nil
}
