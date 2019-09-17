package Dao

import (
	"./Cache"
	"./Database"
)

//主机模式下的缓存初始化：清除缓存后将数据库中的数据写一部分到缓存中
//
//初始化过程中有任何出错都返回error，顺利完成就返回nil
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
