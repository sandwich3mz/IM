package db

import (
	"IM/configs"
	"IM/internel/db/ent"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"time"
)

var url = ""
var DB *ent.Client

func init() {
	dbConf := configs.Conf.DBConfig
	url = fmt.Sprintf("postgres://%s:%s@%s:%v/%s?sslmode=disable&TimeZone=Asia/Shanghai", dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database)
	fmt.Printf("db_url is :%s\n", url)
	DB = NewDBClient(dbConf)
}

func NewDBClient(dbConf configs.DBConfig) *ent.Client {
	dataSourceName := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", dbConf.Host, dbConf.Port, dbConf.Username, dbConf.Password, dbConf.Database)
	logrus.Debugf("dsn: %s\n", dataSourceName)
	db, err := sql.Open("pgx", dataSourceName)
	if err != nil {
		panic(fmt.Sprintf("new db client failed: %v", err))
	}
	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetConnMaxIdleTime(time.Minute * 10)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	if err != nil {
		logrus.Errorf("failed at creating ent client with db %s, err: %v", dataSourceName, err)
		panic(fmt.Sprintf("new db client failed: %v", err))
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}
