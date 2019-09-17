package Cache

import "github.com/garyburd/redigo/redis"

var db redis.Conn = nil
var EX = 6000   //设置每个缓存的存活时间，单位秒
var ranEX = 300 //设置缓存的随机变化范围，单位秒。防止缓存雪崩
var InitN = 300 //设置初始缓存数量，在服务器启动时写入多少记录，单位条

func RedisConnect() (redis.Conn, error) {
	if db != nil {
		return db, nil
	}
	db, err := redis.Dial("tcp", "127.0.0.1:6379")
	return db, err
}
