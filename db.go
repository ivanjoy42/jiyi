package main

import (
	"strings"

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
	sql := `INSERT INTO card(front, back, kind) VALUES(? ,? ,1)`
	db.Exec(sql, front, back)
}

func selectCard() (res []Card) {
	sql := `SELECT * FROM card`
	db.Select(&res, sql)
	return
}

func getCard(cardId string) (res Card) {
	sql := `SELECT * FROM card WHERE card_id=?`
	db.Get(&res, sql, cardId)
	return
}

func delCard(cardId string) {
	sql := `DELETE FROM card WHERE card_id=?`
	db.Exec(sql, cardId)
}

func updateCard(cardId, front, back string) {
	sql := `UPDATE card SET front=?,back=? WHERE card_id=?`
	db.Exec(sql, front, back, cardId)
}

type Deck struct {
	DeckId   int
	Kind     int
	DeckName string
}

func selectDeck() (res []Deck) {
	sql := `SELECT * FROM deck`
	db.Select(&res, sql)
	return
}

func insertDeck(deck string) {
	sql := `INSERT INTO deck(kind, deck_name) VALUES(1, ?)`
	db.Exec(sql, deck)
}

func getDeck(DeckId string) (res Deck) {
	sql := `SELECT * FROM deck WHERE deck_id=?`
	db.Get(&res, sql, DeckId)
	return
}

func selectCardByDeckId(DeckId string) (res []Card) {
	sql := `SELECT card.* FROM card, deck, card_deck WHERE 
		deck.deck_id=? AND 
		card.card_id=card_deck.card_id AND 
		deck.deck_id=card_deck.deck_id`
	db.Select(&res, sql, DeckId)
	return
}

func updateDeck(deckId, deckName string) {
	sql := `UPDATE deck SET deck_name=? WHERE deck_id=?`
	db.Exec(sql, deckName, deckId)
}

// 修改卡片与卡组的关联
//
// 1.删除旧的关联；
// 2.根据卡片的front获取卡片ID；
// 3.添加新的关联。
func updateCardDeck(deckId, kind, str string) {
	deleteCardDeckByDeckId(deckId)

	str = strings.TrimSpace(str)
	str = strings.ReplaceAll(str, "\r", "\n")
	str = strings.ReplaceAll(str, " ", "\n")
	front := strings.Split(str, "\n")
	cards := []Card{}
	sql := `SELECT * FROM card WHERE kind=? AND front IN(?)`
	sql, args, _ := sqlx.In(sql, kind, front)
	db.Select(&cards, sql, args...)

	insertCardDeck(cards, deckId)
}

func deleteCardDeckByDeckId(deckId string) {
	sql := `DELETE FROM card_deck WHERE deck_id=?`
	db.Exec(sql, deckId)
}

func insertCardDeck(cards []Card, deckId string) {
	cardDeck := []map[string]interface{}{}
	for _, v := range cards {
		cardId := v.CardId
		row := map[string]interface{}{"cardId": cardId, "deckId": deckId}
		cardDeck = append(cardDeck, row)
	}
	db.NamedExec(`INSERT INTO card_deck (card_id, deck_id) 
		VALUES (:cardId, :deckId)`, cardDeck)
}
