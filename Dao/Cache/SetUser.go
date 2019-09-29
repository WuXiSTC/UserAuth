package Cache

import "math/rand"

func SetUser(ID, PASS string) (interface{}, error) {
	db, err := RedisConnect()
	if err != nil {
		return nil, err
	}
	if Conf.RandExistTime <= 0 {
		return db.Do("SET", ID, PASS)
	}
	return db.Do("SET", ID, PASS, "EX", Conf.ExistTime+uint64(rand.Int63n(Conf.RandExistTime)))
}
