package Dao

import (
	"./Cache"
)

type SlaveConfig struct {
	SlaveMode  bool   `yaml:"SlaveMode"`  //是否是从属模式
	MasterHost string `yaml:"MasterHost"` //如果是从属模式，则在此指定主服务的Redis地址
	MasterPort string `yaml:"MasterPort"` //如果是从属模式，则在此指定主服务的Redis地址
}

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
