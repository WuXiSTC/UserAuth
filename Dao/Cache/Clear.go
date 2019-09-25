package Cache

func Flushall() (interface{}, error) {
	db, err := RedisConnect()
	if err != nil {
		return nil, err
	}
	return db.Do("FLUSHALL")
}
