package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:123456@/jiyi")
}

type Card struct {
	CardId   int
	Front    string
	Back     string
	Category int
}

func cardInsert(front, back string) {
	sql := "INSERT INTO card(front, back, category) VALUES(? ,? ,1)"
	db.Exec(sql, front, back)
}

func getCard() []Card {
	sql := "SELECT * FROM card"
	query, _ := db.Query(sql)
	var res []Card
	var row Card
	for query.Next() {
		query.Scan(&row.CardId, &row.Front, &row.Back, &row.Category)
		res = append(res, row)
	}
	return res
}

func getCardById(cardId string) Card {
	sql := "SELECT * FROM card WHERE cardId=?"
	var row Card
	db.QueryRow(sql, cardId).Scan(&row.CardId, &row.Front, &row.Back, &row.Category)
	return row
}
