package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/iancoleman/strcase"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	db, _ = sqlx.Connect("mysql", "root:123456@/jiyi")
	db.MapperFunc(strcase.ToSnake)
}

type Card struct {
	CardId int
	Kind   int
	Front  string
	Back   string
}

func insertCard(front, back string) {
	sql := "INSERT INTO card(front, back, kind) VALUES(? ,? ,1)"
	db.Exec(sql, front, back)
}

func selectCard() (res []Card) {
	sql := "SELECT * FROM card"
	db.Select(&res, sql)
	return
}

func getCard(cardId string) (res Card) {
	sql := "SELECT * FROM card WHERE card_id=?"
	db.Get(&res, sql, cardId)
	return
}

func delCard(cardId string) {
	sql := "DELETE FROM card WHERE card_id=?"
	db.Exec(sql, cardId)
}

func updateCard(cardId, front, back string) {
	sql := "UPDATE card SET front=?,back=? WHERE card_id=?"
	db.Exec(sql, front, back, cardId)
}

type Deck struct {
	DeckId   int
	DeckName string
}

func selectDeck() (res []Deck) {
	sql := "SELECT * FROM deck"
	db.Select(&res, sql)
	return
}

func insertDeck(deck string) {
	sql := "INSERT INTO deck(deck_name) VALUES(?)"
	db.Exec(sql, deck)
}

func getDeck(DeckId string) (res Deck) {
	sql := "SELECT * FROM deck WHERE deck_id=?"
	db.Get(&res, sql, DeckId)
	return
}

func selectCardByDeckId(DeckId string) (res []Card) {
	sql := "SELECT card.* FROM card, deck, card_deck WHERE " +
		"deck.deck_id=? AND " +
		"card.card_id=card_deck.card_id AND " +
		"deck.deck_id=card_deck.deck_id"
	db.Select(&res, sql, DeckId)
	return
}
