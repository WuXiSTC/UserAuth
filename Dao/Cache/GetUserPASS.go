package Cache

import "github.com/garyburd/redigo/redis"

/*获取用户密码*/
func GetUserPASS(ID string) (string, error) {
	db, err := RedisConnect()
	if err != nil {
		return "", err
	}
	return redis.String(db.Do("GET", ID))
}
