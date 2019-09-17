package Cache

func DelUser(ID string) (interface{}, error) {
	db, err := RedisConnect()
	if err != nil {
		return nil, err
	}
	return db.Do("DEL", ID)
}
