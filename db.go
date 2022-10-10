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

// 卡片
type Card struct {
	CardId int
	KindId int
	Front  string
	Back   string
	Helper string
	Pinyin string
}

func (card *Card) get(cardId int) *Card {
	sql := `SELECT * FROM card WHERE card_id=?`
	db.Get(card, sql, cardId)
	return card
}

// todo：分页
func (card *Card) list(kindId int) (res []Card) {
	sql := `SELECT * FROM card WHERE kind_id=? LIMIT 100`
	db.Select(&res, sql, kindId)
	return
}

func (card *Card) search(kindId int, query string) (res []Card) {
	fronts := splitSpace(query)
	sql := `SELECT * FROM card WHERE kind_id=? AND front IN(?)`
	sql, args, _ := sqlx.In(sql, kindId, fronts)
	db.Select(&res, sql, args...)
	return
}

func (card *Card) selectCardIds(kindId int, front string) (res []Card) {
	frontArray := splitSpace(front)
	sql := `SELECT * FROM card WHERE kind_id=? AND front IN(?) ORDER BY FIELD(front, ?)`
	sql, args, _ := sqlx.In(sql, kindId, frontArray, frontArray)
	db.Select(&res, sql, args...)
	return
}

func (card *Card) selectByDeckId(deckId int) (res []Card) {
	sql := `SELECT card.* FROM card, deck, card_deck 
			WHERE deck.deck_id=? 
			AND card.card_id=card_deck.card_id 
			AND deck.deck_id=card_deck.deck_id
			ORDER BY card_deck_id`
	db.Select(&res, sql, deckId)
	return
}

func (card *Card) insert() {
	sql := `INSERT INTO card(kind_id, front, back, helper, pinyin) VALUES(? ,? ,?, ?, ?)`
	db.Exec(sql, card.KindId, card.Front, card.Back, card.Helper, card.Pinyin)
}

func (card *Card) update() {
	sql := `UPDATE card SET front=?, back=?, helper=?, pinyin=? WHERE card_id=?`
	db.Exec(sql, card.Front, card.Back, card.Helper, card.Pinyin, card.CardId)
}

// 删除卡片（todo：事物）
//
// 1.删除卡片
// 2.删除卡片与卡组的关联
func (card *Card) deleteTx(cardId int) {
	card.delete(cardId)
	cardDeck.deleteByCardId(cardId)
}

func (card *Card) delete(cardId int) {
	sql := `DELETE FROM card WHERE card_id=?`
	db.Exec(sql, cardId)
}

// 卡组
type Deck struct {
	DeckId   int
	KindId   int
	DeckName string
}

func (deck *Deck) get(deckId int) *Deck {
	sql := `SELECT * FROM deck WHERE deck_id=?`
	db.Get(deck, sql, deckId)
	return deck
}

func (deck *Deck) list(kindId int) (res []Deck) {
	sql := `SELECT * FROM deck WHERE kind_id=? LIMIT 100`
	db.Select(&res, sql, kindId)
	return
}

// 添加卡组（todo：事物）
//
// 1.添加卡组
// 2.获取卡片ID
// 3.添加卡片与卡组的关联
func (deck *Deck) insertTx(fronts string) {
	deckId := deck.insert()
	card := Card{}
	cardArray := card.selectCardIds(deck.KindId, fronts)
	cardDeck.insert(cardArray, deckId)
}

func (deck *Deck) insert() int {
	sql := `INSERT INTO deck(deck_name, kind_id) VALUES(?, ?)`
	res, _ := db.Exec(sql, deck.DeckName, deck.KindId)
	lastId64, _ := res.LastInsertId()
	return int(lastId64)
}

// 更新卡组（todo：事物）
//
// 1.更新卡组
// 2.删除旧的关联
// 3.获取卡片ID
// 4.添加卡片与卡组的关联
func (deck *Deck) updateTx(fronts string) {
	deck.update()
	cardDeck.deleteByDeckId(deck.DeckId)
	card := Card{}
	cardArray := card.selectCardIds(deck.KindId, fronts)
	cardDeck.insert(cardArray, deck.DeckId)
}

func (deck *Deck) update() {
	sql := `UPDATE deck SET deck_name=? WHERE deck_id=?`
	db.Exec(sql, deck.DeckName, deck.DeckId)
}

// 删除卡组（todo：事物）
//
// 1.删除卡组
// 2.删除卡片与卡组的关联
func (deck *Deck) deleteTx(deckId int) {
	deck.delete(deckId)
	cardDeck.deleteByDeckId(deckId)
}

func (deck *Deck) delete(deckId int) {
	sql := `DELETE FROM deck WHERE deck_id=?`
	db.Exec(sql, deckId)
}

// 卡片卡组关联
type CardDeck struct{}

var cardDeck CardDeck

// 卡片与卡组的关联操作
func (cardDeck *CardDeck) insert(cardArray []Card, deckId int) {
	cardDeckArray := []map[string]interface{}{}
	for _, v := range cardArray {
		cardId := v.CardId
		row := map[string]interface{}{"cardId": cardId, "deckId": deckId}
		cardDeckArray = append(cardDeckArray, row)
	}
	sql := `INSERT INTO card_deck (card_id, deck_id) VALUES (:cardId, :deckId)`
	db.NamedExec(sql, cardDeckArray)
}

func (cardDeck *CardDeck) deleteByCardId(cardId int) {
	sql := `DELETE FROM card_deck WHERE card_id=?`
	db.Exec(sql, cardId)
}

func (cardDeck *CardDeck) deleteByDeckId(deckId int) {
	sql := `DELETE FROM card_deck WHERE deck_id=?`
	db.Exec(sql, deckId)
}

// 类型
type Kind struct {
	KindId   int
	KindName string
}

func (kind *Kind) get(kindId int) *Kind {
	kind.KindId = kindId
	switch kindId {
	case 1:
		kind.KindName = "普通"
	case 2:
		kind.KindName = "汉字"
	case 3:
		kind.KindName = "词语"
	case 4:
		kind.KindName = "古诗文"
	}
	return kind
}

// 公用函数
func splitSpace(s string) (res []string) {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\r", " ")
	s = strings.ReplaceAll(s, "\n", " ")
	res = strings.Fields(s)
	return
}
