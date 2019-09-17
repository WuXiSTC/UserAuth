package Cache

func Clear() (interface{}, error) {
	db, err := RedisConnect()
	if err != nil {
		return nil, err
	}
	return db.Do("CLEAR")
}
