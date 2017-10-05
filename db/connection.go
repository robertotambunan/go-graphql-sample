package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DbUsername = "root"
	DbPassword = "root"
	DbName     = "belajar"
)

func OpenConnection() *sql.DB {
	conn := fmt.Sprintf("%s:%s@/%s?charset=utf8", DbUsername, DbPassword, DbName)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Println("Cannot Open Connection MySQL Database")
		return nil
	}
	return db
}

func CloseConnection(db *sql.DB) {
	db.Close()
}
