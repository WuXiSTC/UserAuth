package Cache

import (
	"../../util"
	"github.com/garyburd/redigo/redis"
)

type config struct {
	ExistTime     uint64 `yaml:"ExistTime"`     //设置每个缓存的存活时间，单位秒
	RandExistTime int64  `yaml:"RandExistTime"` //设置缓存的随机变化范围，单位秒。防止缓存雪崩
	InitialCache  uint64 `yaml:"InitialCache"`  //设置初始缓存数量，在服务器启动时写入多少记录，单位条
	Network       string `yaml:"Network"`
	Address       string `yaml:"Address"`
}

var db redis.Conn = nil
var Conf = config{6000, 300, 300, "tcp", "127.0.0.1:6379"}

func RedisConnect() (redis.Conn, error) {
	if db != nil {
		return db, nil
	}
	util.GetConf("CacheConfig.yaml", &Conf)
	db, err := redis.Dial(Conf.Network, Conf.Address)
	return db, err
}
