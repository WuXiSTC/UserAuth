package Database

import (
	"../../util"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	Driver       string `yaml:"Driver"`
	DataSource   string `yaml:"DataSource"`   //数据库连接
	MaxOpenConns int    `yaml:"MaxOpenConns"` //数据库中同时存在的连接总数
}

var db *sql.DB = nil
var Conf = config{"mysql", "WXSTC:WXSTC@/WXSTC", 2}

func RDBConnect() (*sql.DB, error) {
	if db != nil {
		return db, nil
	}
	util.GetConf("DatabaseConfig.yaml", &Conf)
	db, err := sql.Open(Conf.Driver, Conf.DataSource)
	if err != nil {
		return db, err
	}
	db.SetMaxOpenConns(Conf.MaxOpenConns)
	return db, nil
}
