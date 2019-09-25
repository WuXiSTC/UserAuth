package Dao

import (
	"./Cache"
	"./Database"
)

//主机模式下的缓存初始化：清除缓存后将数据库中的数据写一部分到缓存中
//
//初始化过程中有任何出错都返回error，顺利完成就返回nil
func CacheInit() error {
	_, _ = Cache.Flushall()
	rows, err := Database.GetALL(Cache.Conf.InitialCache)
	if err != nil {
		return err
	}
	for ID, PASS := range rows {
		_, _ = Cache.SetUser(ID, PASS)
	}
	return nil
}
