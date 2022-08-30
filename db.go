package main

import (
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	db, _ = sqlx.Connect("mysql", "root:123456@/jiyi")
	db.MapperFunc(camelCase)
}

type Card struct {
	CardId   int
	Front    string
	Back     string
	Category int
}

func camelCase(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
}

func insertCard(front, back string) {
	sql := "INSERT INTO card(front, back, category) VALUES(? ,? ,1)"
	db.Exec(sql, front, back)
}

func selectCard() (res []Card) {
	sql := "SELECT * FROM card"
	db.Select(&res, sql)
	return
}

func getCard(cardId string) (res Card) {
	sql := "SELECT * FROM card WHERE cardId=?"
	db.Get(&res, sql, cardId)
	return
}

func delCard(cardId string) {
	sql := "DELETE FROM card WHERE cardId=?"
	db.Exec(sql, cardId)
}

func updateCard(cardId, front, back string) {
	sql := "UPDATE card SET front=?,back=? WHERE cardId=?"
	db.Exec(sql, front, back, cardId)
}
