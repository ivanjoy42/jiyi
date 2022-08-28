package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:123456@/jiyi")
}

func cardInsert(front, back string) {
	sql := "INSERT INTO card(front, back, type) VALUES(? ,? ,1)"
	db.Exec(sql, front, back)
}
