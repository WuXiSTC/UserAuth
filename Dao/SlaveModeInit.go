package Dao

import (
	"../util"
	"./Cache"
)

//从机模式的配置信息
type SlaveConfig struct {
	SlaveMode     bool   `yaml:"SlaveMode"`    //是否是从属模式
	SlaveAddress  string `yaml:"SlaveAddress"` //如果是从属模式，则在此指定主服务的Redis地址
	SlaveInit     bool   `yaml:"SlaveInit"`    //是否运行slaveof指令进行从机初始化
	MasterHost    string `yaml:"MasterHost"`   //如果是从属模式，则在此指定主服务的Redis地址
	MasterPort    uint16 `yaml:"MasterPort"`   //如果是从属模式，则在此指定主服务的Redis端口
	MasterAppAddr string `yaml:"MasterAppAddr"`
}

var SlaveConf = SlaveConfig{
	false,
	"127.0.0.1:6379",
	false,
	"redis",
	6379,
	"127.0.0.1:80",
}

func SlaveConfigure(path string) {
	util.GetConf(path, &SlaveConf) //先读配置文件
}

//从机模式下的缓存初始化：用slaveof指令将Redis设置为指向主机的从机，主机的指定见配置文件
//
//初始化过程中有任何出错都返回error，顺利完成就返回nil
func SlaveModeInit() error {
	Cache.Conf.Address = SlaveConf.SlaveAddress
	db, err := Cache.RedisConnect()
	if err != nil {
		return err
	}
	if SlaveConf.SlaveInit && db != nil {
		_, err = db.Do("slaveof", SlaveConf.MasterHost, SlaveConf.MasterPort)
		return err
	}
	return nil
}
