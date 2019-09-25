package Cache

import "math/rand"

func SetUser(ID, PASS string) (interface{}, error) {
	db, err := RedisConnect()
	if err != nil {
		return nil, err
	}
	return db.Do("SET", ID, PASS, "EX", Conf.ExistTime+uint64(rand.Int63n(Conf.RandExistTime)))
}
