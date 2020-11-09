package util

import (
	"database/sql"
	"os"
)

// DBConn Connect data
func DBConn() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_GOLANG_LEARN_02"))
	PanicError(err)

	return db
}
