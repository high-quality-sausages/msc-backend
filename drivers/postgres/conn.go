package postgres

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/high-quality-sausages/msc-backend/conf"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var config = conf.GetConfig()

	var pgInfo string

	fmt.Println(config.PgConf.Host)

	if config.PgConf.Password == "" {
		pgInfo = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
			config.PgConf.Host, config.PgConf.Port, config.PgConf.User, config.PgConf.DBName)
	} else {
		pgInfo = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			config.PgConf.Host, config.PgConf.Port, config.PgConf.User, config.PgConf.DBName, config.PgConf.Password)
	}

	db, _ = sql.Open("postgres", pgInfo)
	db.SetMaxIdleConns(1000)

	err := db.Ping()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func DBConn() *sql.DB {
	return db
}
