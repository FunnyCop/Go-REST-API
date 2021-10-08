package config

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Connect to MySQL
func MysqlConnect() *sql.DB {

	db, err := sql.Open("mysql", "root:68781020lbP!@/job_applications")

	if err != nil {

		panic(err.Error())

	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db

}
