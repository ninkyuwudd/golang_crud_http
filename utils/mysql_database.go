package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func MysqlConnection() *sql.DB {

	if db != nil {
		return db
	}
	//db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/golang_learn")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_learn")
	if err != nil {
		panic(err)
	}
	return db

}
