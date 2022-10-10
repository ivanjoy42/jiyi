package main

import (
	"strconv"
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

func (card *Card) get() *Card {
	sql := `SELECT * FROM card WHERE card_id=?`
	db.Get(card, sql, card.CardId)
	return card
}

// todo：分页
func (card *Card) getAll() (cards []Card) {
	sql := `SELECT * FROM card WHERE kind_id=? LIMIT 100`
	db.Select(&cards, sql, card.KindId)
	return
}

func (card *Card) search(query string) (cards []Card) {
	fronts := splitSpace(query)
	sql := `SELECT * FROM card WHERE kind_id=? AND front IN(?)`
	sql, args, _ := sqlx.In(sql, card.KindId, fronts)
	db.Select(&cards, sql, args...)
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
func (card *Card) deleteTx() {
	card.delete()
	cardDeck.deleteByCardId(card.CardId)
}

func (card *Card) delete() {
	sql := `DELETE FROM card WHERE card_id=?`
	db.Exec(sql, card.CardId)
}

// 卡组
type Deck struct {
	DeckId   int
	KindId   int
	DeckName string
}

func getDeck(DeckId string) (res Deck) {
	sql := `SELECT * FROM deck WHERE deck_id=?`
	db.Get(&res, sql, DeckId)
	return
}

func selectDeck(kindId string) (res []Deck) {
	sql := `SELECT * FROM deck WHERE kind_id=? LIMIT 100`
	db.Select(&res, sql, kindId)
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

func selectCardIds(kindId, front string) (cards []Card) {
	frontArray := splitSpace(front)
	sql := `SELECT * FROM card WHERE kind_id=? AND front IN(?) ORDER BY FIELD(front, ?)`
	sql, args, _ := sqlx.In(sql, kindId, frontArray, frontArray)
	db.Select(&cards, sql, args...)
	return
}

// 添加卡组（todo：事物）
//
// 1.添加卡组
// 2.获取卡片ID
// 3.添加卡片与卡组的关联
func insertDeckTxn(deckName, kindId, cards string) {
	deckId := insertDeck(deckName, kindId)
	cardIds := selectCardIds(kindId, cards)
	insertCardDeck(cardIds, deckId)
}

func insertDeck(deckName, kindId string) (id string) {
	sql := `INSERT INTO deck(deck_name, kind_id) VALUES(?, ?)`
	res, _ := db.Exec(sql, deckName, kindId)
	lastId, _ := res.LastInsertId()
	id = strconv.FormatInt(lastId, 10)
	return
}

// 更新卡组（todo：事物）
//
// 1.更新卡组
// 2.删除旧的关联
// 3.获取卡片ID
// 4.添加卡片与卡组的关联
func updateDeckTxn(deckId, deckName, kindId, cards string) {
	updateDeck(deckId, deckName)
	deleteCardDeckByDeckId(deckId)
	cardIds := selectCardIds(kindId, cards)
	insertCardDeck(cardIds, deckId)
}

func updateDeck(deckId, deckName string) {
	sql := `UPDATE deck SET deck_name=? WHERE deck_id=?`
	db.Exec(sql, deckName, deckId)
}

func deleteDeck(deckId string) {
	sql := `DELETE FROM deck WHERE deck_id=?`
	db.Exec(sql, deckId)
	deleteCardDeckByDeckId(deckId)
}

// 卡片与卡组的关联操作
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

// 卡片卡组关联
type CardDeck struct{}

var cardDeck CardDeck

func (cardDeck *CardDeck) deleteByCardId(cardId int) {
	sql := `DELETE FROM card_deck WHERE card_id=?`
	db.Exec(sql, cardId)
}

func deleteCardDeckByDeckId(deckId string) {
	sql := `DELETE FROM card_deck WHERE deck_id=?`
	db.Exec(sql, deckId)
}

// 类型
type Kind struct {
	KindId   int
	KindName string
}

func getKind(kindId string) (res Kind) {
	switch kindId {
	case "1":
		res = Kind{1, "基本"}
	case "2":
		res = Kind{2, "汉字"}
	case "3":
		res = Kind{3, "词语"}
	case "4":
		res = Kind{4, "古诗文"}
	}
	return
}

func (kind *Kind) get() *Kind {
	switch kind.KindId {
	case 1:
		kind.KindId = 1
		kind.KindName = "普通"
	case 2:
		kind.KindId = 2
		kind.KindName = "汉字"
	case 3:
		kind.KindId = 3
		kind.KindName = "词语"
	case 4:
		kind.KindId = 4
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
