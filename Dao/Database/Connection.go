package Database

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var db *sql.DB = nil

func RDBConnect() (*sql.DB, error) {
	if db != nil {
		return db, nil
	}
	db, err := sql.Open("mysql", "WXSTC:WXSTC@/WXSTC") //数据库连接
	if err != nil {
		return db, err
	}
	db.SetMaxOpenConns(4) //数据库中同时存在的连接总数
	return db, nil
}

