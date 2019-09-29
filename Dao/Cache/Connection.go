package Cache

import (
	"../../util"
	"github.com/garyburd/redigo/redis"
)

type Config struct {
	ExistTime     uint64 `yaml:"ExistTime"`     //设置每个缓存的存活时间，单位秒
	RandExistTime int64  `yaml:"RandExistTime"` //设置缓存的随机变化范围，单位秒。防止缓存雪崩
	Network       string `yaml:"Network"`
	Address       string `yaml:"Address"`
	InitAtStart   bool   `yaml:"InitAtStart"`
	InitialCache  uint64 `yaml:"InitialCache"` //设置初始缓存数量，在服务器启动时写入多少记录，单位条
	MeetAtStart   bool   `yaml:"MeetAtStart"`  //是否在启动时meet某个redis
	MeetAddress   string `yaml:"MeetAddress"`  //redis-1 #要meet的地址
	MeetPort      uint16 `yaml:"MeetPort"`     //要meet的端口
}

var db redis.Conn = nil
var Conf = Config{
	6000, 300,
	"tcp", "127.0.0.1:6379",
	true, 300,
	false, "127.0.0.1", 6379,
}

func ConfigureRedis(path string) {
	util.GetConf(path, &Conf)
}

func RedisConnect() (redis.Conn, error) {
	if db != nil {
		return db, nil
	}
	db, err := redis.Dial(Conf.Network, Conf.Address)
	return db, err
}
