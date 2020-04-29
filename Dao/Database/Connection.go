package Database

import (
	"database/sql"
	"fmt"
	"gitee.com/WuXiSTC/UserAuth/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocql/gocql"
)

type RDBConfig struct {
	Driver       string `yaml:"Driver"`
	Database     string `yaml:"Database"`
	Username     string `yaml:"Username"`
	Password     string `yaml:"Password"`
	Host         string `yaml:"Host"`
	Port         uint16 `yaml:"Port"`
	ConfString   string `yaml:"ConfString"`
	MaxOpenConns int    `yaml:"MaxOpenConns"`
}

type CDBConfig struct {
	KeySpace string   `yaml:"KeySpace"`
	Hosts    []string `yaml:"Hosts"`
	Port     uint16   `yaml:"Port"`
}

type Config struct {
	UseCassandra bool      `yaml:"UseCassandra"`
	RDBConf      RDBConfig `yaml:"RDBConfig"`
	CDBConf      CDBConfig `yaml:"CDBConfig"`
}

var db *sql.DB = nil
var Conf = Config{
	false,
	RDBConfig{
		"mysql", "WXSTC", "WXSTC", "WXSTC",
		"127.0.0.1", 3306, "charset=utf8", 2,
	},
	CDBConfig{
		"wxstc", []string{"127.0.0.1"}, 9042,
	},
}

func ConfigureDatabase(path string) {
	util.GetConf(path, &Conf)
}

func RDBConnect() (*sql.DB, error) {
	if db != nil {
		return db, nil
	}
	conf := Conf.RDBConf
	DataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		conf.Username, conf.Password, conf.Host, conf.Port, conf.Database, conf.ConfString)
	db, err := sql.Open(conf.Driver, DataSource)
	if err != nil {
		return db, err
	}
	db.SetMaxOpenConns(conf.MaxOpenConns)
	return db, nil
}

var cluster *gocql.ClusterConfig = nil

func CDBConnect() (*gocql.ClusterConfig, error) {
	if cluster != nil {
		return cluster, nil
	}
	conf := Conf.CDBConf
	cluster = gocql.NewCluster(conf.Hosts...)
	cluster.Keyspace = conf.KeySpace
	return cluster, nil
}
