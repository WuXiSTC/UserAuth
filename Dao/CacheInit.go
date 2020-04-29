package Dao

import (
	"gitee.com/WuXiSTC/UserAuth/Dao/Cache"
	"gitee.com/WuXiSTC/UserAuth/Dao/Database"
	"gitee.com/WuXiSTC/UserAuth/util"
)

//主机模式下的缓存初始化：清除缓存后将数据库中的数据写一部分到缓存中
//
//初始化过程中有任何出错都返回error，顺利完成就返回nil
func CacheInit() error {
	if Cache.Conf.MeetAtStart {
		_, _ = Cache.Flushall()
		rows, err := Database.GetALL(Cache.Conf.InitialCache)
		util.LogE(err)
		for ID, PASS := range rows {
			_, err = Cache.SetUser(ID, PASS)
			util.LogE(err)
		}
	} /*TODO:集群接口暂不支持
	if Cache.Conf.MeetAtStart {
		db, err := Cache.RedisConnect()
		util.LogE(err)
		if db != nil {
			_, err = db.Do("CLUSTER", "MEET", Cache.Conf.MeetAddress, Cache.Conf.MeetPort)
			util.LogE(err)
		}
	}*/
	return nil
}
