package Dao

import (
	"./Cache"
)

//从机模式的配置信息
type SlaveConfig struct {
	SlaveMode  bool   `yaml:"SlaveMode"`  //是否是从属模式
	MasterHost string `yaml:"MasterHost"` //如果是从属模式，则在此指定主服务的Redis地址
	MasterPort string `yaml:"MasterPort"` //如果是从属模式，则在此指定主服务的Redis地址
}

//从机模式下的缓存初始化：用slaveof指令将Redis设置为指向主机的从机，主机的指定见配置文件
//
//初始化过程中有任何出错都返回error，顺利完成就返回nil
func SlaveModeInit(conf SlaveConfig) error {
	db, err := Cache.RedisConnect()
	if err != nil {
		return err
	}
	if db != nil {
		_, err = db.Do("slaveof", conf.MasterHost, conf.MasterPort)
		return err
	}
	return nil
}
