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

// 卡片数据结构
type Card struct {
	CardId int
	Kind   int
	Front  string
	Back   string
	Helper string
	Pinyin string
}

func insertCard(kind, front, back, helper, pinyin string) {
	sql := `INSERT INTO card(kind, front, back, helper, pinyin) VALUES(? ,? ,?, ?, ?)`
	db.Exec(sql, kind, front, back, helper, pinyin)
}

// todo：分页
func selectCard(kind string) (res []Card) {
	sql := `SELECT * FROM card WHERE kind=? LIMIT 100`
	db.Select(&res, sql, kind)
	return
}

func getCard(cardId string) (res Card) {
	sql := `SELECT * FROM card WHERE card_id=?`
	db.Get(&res, sql, cardId)
	return
}

func deleteCard(cardId string) {
	sql := `DELETE FROM card WHERE card_id=?`
	db.Exec(sql, cardId)
}

func updateCard(cardId, cardName, back, helper, pinyin string) {
	sql := `UPDATE card SET front=?, back=?, helper=?, pinyin=? WHERE card_id=?`
	db.Exec(sql, cardName, back, helper, pinyin, cardId)
}

func searchCard(sCardName string) (res []Card) {
	cardName := splitSpace(sCardName)
	sql := `SELECT * FROM card WHERE front IN(?)`
	sql, args, _ := sqlx.In(sql, cardName)
	db.Select(&res, sql, args...)
	return
}

// 卡组数据结构
type Deck struct {
	DeckId   int
	Kind     int
	DeckName string
}

func selectDeck(kind string) (res []Deck) {
	sql := `SELECT * FROM deck WHERE kind=? LIMIT 100`
	db.Select(&res, sql, kind)
	return
}

func insertDeck(deckName, kind string) {
	sql := `INSERT INTO deck(deck_name, kind) VALUES(?, ?)`
	db.Exec(sql, deckName, kind)
}

func getDeck(DeckId string) (res Deck) {
	sql := `SELECT * FROM deck WHERE deck_id=?`
	db.Get(&res, sql, DeckId)
	return
}

func selectCardByDeckId(DeckId string) (res []Card) {
	sql := `SELECT card.* FROM card, deck, card_deck 
			WHERE deck.deck_id=? 
			AND card.card_id=card_deck.card_id 
			AND deck.deck_id=card_deck.deck_id
			ORDER BY card_deck_id`
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
func updateCardDeck(deckId, kind, sFront string) {
	deleteCardDeckByDeckId(deckId)

	front := splitSpace(sFront)
	cards := []Card{}
	sql := `SELECT * FROM card WHERE kind=? AND front IN(?) ORDER BY FIELD(front, ?)`
	sql, args, _ := sqlx.In(sql, kind, front, front)
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
	sql := `INSERT INTO card_deck (card_id, deck_id) VALUES (:cardId, :deckId)`
	db.NamedExec(sql, cardDeck)
}

func deleteDeck(deckId string) {
	sql := `DELETE FROM deck WHERE deck_id=?`
	db.Exec(sql, deckId)

	deleteCardDeckByDeckId(deckId)
}

func splitSpace(s string) (res []string) {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\r", " ")
	s = strings.ReplaceAll(s, "\n", " ")
	res = strings.Fields(s)
	return
}

func getKindName(kind string) (res string) {
	switch kind {
	case "1":
		res = "通用"
	case "2":
		res = "字词"
	}
	return
}
