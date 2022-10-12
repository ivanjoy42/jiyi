package main

import (
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/iancoleman/strcase"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// 公共函数
func splitSpace(s string) (res []string) {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\r", " ")
	s = strings.ReplaceAll(s, "\n", " ")
	res = strings.Fields(s)
	return
}

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

func (c *Card) get(cardId int) *Card {
	sql := `SELECT * FROM card WHERE card_id=?`
	db.Get(c, sql, cardId)
	return c
}

func (c *Card) insert() {
	sql := `INSERT INTO card(kind_id, front, back, helper, pinyin) VALUES(? ,? ,?, ?, ?)`
	db.Exec(sql, c.KindId, c.Front, c.Back, c.Helper, c.Pinyin)
}

func (c *Card) update() {
	sql := `UPDATE card SET front=?, back=?, helper=?, pinyin=? WHERE card_id=?`
	db.Exec(sql, c.Front, c.Back, c.Helper, c.Pinyin, c.CardId)
}

func (c *Card) delete(cardId int) {
	sql := `DELETE FROM card WHERE card_id=?`
	db.Exec(sql, cardId)
}

// 删除卡片（todo：事物）
//
// 1.删除卡片
// 2.删除卡片与卡组的关联
func (c *Card) deleteTx(cardId int) {
	c.delete(cardId)
	cardDeck.deleteByCard(cardId)
}

type Cards []Card

// todo：分页
func (c *Cards) list(kindId int) *Cards {
	sql := `SELECT * FROM card WHERE kind_id=? LIMIT 100`
	db.Select(c, sql, kindId)
	return c
}

func (c *Cards) search(kindId int, query string) *Cards {
	frontArray := splitSpace(query)
	sql := `SELECT * FROM card WHERE kind_id=? AND front IN(?)`
	sql, args, _ := sqlx.In(sql, kindId, frontArray)
	db.Select(c, sql, args...)
	return c
}

func (c *Cards) getId(kindId int, fronts string) (res []int) {
	frontArray := splitSpace(fronts)
	sql := `SELECT card_id FROM card WHERE kind_id=? AND front IN(?) ORDER BY FIELD(front, ?)`
	sql, args, _ := sqlx.In(sql, kindId, frontArray, frontArray)
	db.Select(&res, sql, args...)
	return
}

// 卡组
type Deck struct {
	DeckId   int
	KindId   int
	DeckName string
}

func (d *Deck) get(deckId int) *Deck {
	sql := `SELECT * FROM deck WHERE deck_id=?`
	db.Get(d, sql, deckId)
	return d
}

func (d *Deck) insert() int {
	sql := `INSERT INTO deck(deck_name, kind_id) VALUES(?, ?)`
	res, _ := db.Exec(sql, d.DeckName, d.KindId)
	lastId, _ := res.LastInsertId()
	return int(lastId)
}

func (d *Deck) update() {
	sql := `UPDATE deck SET deck_name=? WHERE deck_id=?`
	db.Exec(sql, d.DeckName, d.DeckId)
}

func (d *Deck) delete(deckId int) {
	sql := `DELETE FROM deck WHERE deck_id=?`
	db.Exec(sql, deckId)
}

// 添加卡组（todo：事物）
//
// 1.添加卡组
// 2.获取卡片ID
// 3.添加卡片与卡组的关联
func (d *Deck) insertTx(fronts string) {
	deckId := d.insert()
	cardIds := cards.getId(d.KindId, fronts)
	cardDeck.insert(cardIds, deckId)
}

// 更新卡组（todo：事物）
//
// 1.更新卡组
// 2.删除旧的关联
// 3.获取卡片ID
// 4.添加卡片与卡组的关联
func (d *Deck) updateTx(fronts string) {
	d.update()
	cardDeck.deleteByDeck(d.DeckId)
	cardIds := cards.getId(d.KindId, fronts)
	cardDeck.insert(cardIds, d.DeckId)
}

// 删除卡组（todo：事物）
//
// 1.删除卡组
// 2.删除卡片与卡组的关联
func (d *Deck) deleteTx(deckId int) {
	d.delete(deckId)
	cardDeck.deleteByDeck(deckId)
}

func (d *Deck) getFronts(deckId int) string {
	res := []string{}
	sql := `SELECT card.front FROM card, deck, card_deck
			WHERE deck.deck_id=?
			AND card.card_id=card_deck.card_id
			AND deck.deck_id=card_deck.deck_id
			ORDER BY card_deck_id`
	db.Select(&res, sql, deckId)
	return strings.Join(res, "\n")
}

type Decks []Deck

func (d *Decks) list(kindId int) *Decks {
	sql := `SELECT * FROM deck WHERE kind_id=? LIMIT 100`
	db.Select(d, sql, kindId)
	return d
}

// 卡片卡组关联
type CardDeck struct {
	CardDeckId int
	CardId     int
	DeckId     int
}

// 卡片与卡组的关联操作
func (cd *CardDeck) insert(cardIds []int, deckId int) {
	cardDeckArray := []map[string]interface{}{}
	for _, v := range cardIds {
		row := map[string]interface{}{"cardId": v, "deckId": deckId}
		cardDeckArray = append(cardDeckArray, row)
	}
	sql := `INSERT INTO card_deck (card_id, deck_id) VALUES (:cardId, :deckId)`
	db.NamedExec(sql, cardDeckArray)
}

func (cardDeck *CardDeck) deleteByCard(cardId int) {
	sql := `DELETE FROM card_deck WHERE card_id=?`
	db.Exec(sql, cardId)
}

func (cardDeck *CardDeck) deleteByDeck(deckId int) {
	sql := `DELETE FROM card_deck WHERE deck_id=?`
	db.Exec(sql, deckId)
}

// 卡片类型
type Kind struct {
	KindId   int
	KindName string
}

func (k *Kind) get(kindId int) *Kind {
	k.KindId = kindId
	switch kindId {
	case 1:
		k.KindName = "其他"
	case 2:
		k.KindName = "汉字"
	case 3:
		k.KindName = "词语"
	case 4:
		k.KindName = "古诗文"
	}
	return k
}

// 学习模式
type Mode struct {
	ModeId   int
	ModeName string
}

var mode Mode

func (m *Mode) get(modeId int) *Mode {
	m.ModeId = modeId
	switch modeId {
	case 1:
		m.ModeName = "背诵"
	case 2:
		m.ModeName = "默写"
	case 3:
		m.ModeName = "认识"
	case 4:
		m.ModeName = "听写"
	}
	return m
}

// 学习
type Learn struct {
	LearnId   int
	ModeId    int
	KindId    int
	DeckId    int
	LearnName string
}

func (l *Learn) get(learnId int) Learn {
	sql := `SELECT * FROM learn WHERE learn_id=?`
	db.Get(l, sql, learnId)
	return *l
}

func (l *Learn) list() (res []Learn) {
	sql := `SELECT * FROM learn`
	db.Select(&res, sql)
	return
}

// 学习卡片
type LearnCard struct {
	LearnCardId int
	LearnId     int
	CardId      int
	rank        int
	skip        int
	updateTime  int
	NextTime    int
}
