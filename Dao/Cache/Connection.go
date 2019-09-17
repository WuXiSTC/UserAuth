package Cache

import (
	"../../util"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type config struct {
	ExistTime     int    `yaml:"ExistTime"`     //设置每个缓存的存活时间，单位秒
	RandExistTime int    `yaml:"RandExistTime"` //设置缓存的随机变化范围，单位秒。防止缓存雪崩
	InitialCache  int    `yaml:"InitialCache"`  //设置初始缓存数量，在服务器启动时写入多少记录，单位条
	Network       string `yaml:"Network"`
	Address       string `yaml:"Address"`
}

var db redis.Conn = nil
var Conf = config{6000, 300, 300, "tcp", "127.0.0.1:6379"}

func RedisConnect() (redis.Conn, error) {
	if db != nil {
		return db, nil
	}
	yamlFile, err := ioutil.ReadFile("CacheConfig.yaml")
	util.LogE(err)
	if err == nil {
		err = yaml.Unmarshal(yamlFile, &Conf)
		util.LogE(err)
	}
	db, err := redis.Dial(Conf.Network, Conf.Address)
	return db, err
}
