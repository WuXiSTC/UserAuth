package Daemons

import "../../util"

type Config struct {
	NumOfInsertDaemons uint16 `yaml:"NumOfInsertDaemons"`
	NumOfDeleteDaemons uint16 `yaml:"NumOfDeleteDaemons"`
	NumOfUpdateDaemons uint16 `yaml:"NumOfUpdateDaemons"`
}

var Conf = Config{4, 4, 4}

//按照配置文件启动所有的守护进程
func StartDaemons(ConfPath string) {
	util.GetConf(ConfPath, &Conf)
	for i := uint16(0); i < Conf.NumOfInsertDaemons; i++ {
		go InsertDeamon()
	}
	for i := uint16(0); i < Conf.NumOfUpdateDaemons; i++ {
		go UpdateDaemon()
	}
	for i := uint16(0); i < Conf.NumOfDeleteDaemons; i++ {
		go DeleteDaemon()
	}
}
