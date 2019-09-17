package Cache

import (
	"github.com/garyburd/redigo/redis"
)

func Ping() (string, error) {
	db, err := RedisConnect()
	if err != nil {
		return "", err
	}
	pong, err := redis.String(db.Do("PING"))
	if err != nil {
		return "", err
	}
	return pong, nil
}
