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

func (c *Card) get(cardId int) (res Card) {
	sql := `SELECT * FROM card WHERE card_id=?`
	db.Get(&res, sql, cardId)
	return
}

func (c *Card) insert() int {
	sql := `INSERT INTO card(kind_id, front, back, helper, pinyin) VALUES(? ,? ,?, ?, ?)`
	res, _ := db.Exec(sql, c.KindId, c.Front, c.Back, c.Helper, c.Pinyin)
	lastId, _ := res.LastInsertId()
	return int(lastId)
}

func (c *Card) insertTx(deckId int) {
	cardId := c.insert()
	cardIds := []int{cardId}
	cardDeck.insert(cardIds, deckId)
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

// todo：分页
func (c *Card) list(kindId, deckId int) (res []Card) {
	sql := `SELECT card.* FROM card, card_deck 
			WHERE card.card_id=card_deck.card_id 
			AND kind_id=? 
			AND deck_id=? LIMIT 100`
	db.Select(&res, sql, kindId, deckId)
	return
}

func (c *Card) search(kindId int, query string) (res []Card) {
	frontArray := splitSpace(query)
	sql := `SELECT * FROM card WHERE kind_id=? AND front IN(?)`
	sql, args, _ := sqlx.In(sql, kindId, frontArray)
	db.Select(&res, sql, args...)
	return
}

func (c *Card) getIds(kindId int, fronts string) (res []int) {
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

func (d *Deck) get(deckId int) (res Deck) {
	sql := `SELECT * FROM deck WHERE deck_id=?`
	db.Get(&res, sql, deckId)
	return
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
	cardIds := card.getIds(d.KindId, fronts)
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
	cardIds := card.getIds(d.KindId, fronts)
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

func (d *Deck) list(kindId int) (res []Deck) {
	sql := `SELECT * FROM deck WHERE kind_id=? LIMIT 100`
	db.Select(&res, sql, kindId)
	return res
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

func (k *Kind) get(kindId int) (res Kind) {
	sql := `SELECT * FROM kind WHERE kind_id=?`
	db.Get(&res, sql, kindId)
	return
}

func (k *Kind) insert() {
	sql := `INSERT INTO kind(kind_name) VALUES(?)`
	db.Exec(sql, k.KindName)
}

func (k *Kind) update() {
	sql := `UPDATE kind SET kind_name=? WHERE kind_id=?`
	db.Exec(sql, k.KindName, k.KindId)
}

func (k *Kind) delete(kindId int) {
	sql := `DELETE FROM kind WHERE kind_id=?`
	db.Exec(sql, kindId)
}

func (k *Kind) list() (res []Kind) {
	sql := `SELECT * FROM kind`
	db.Select(&res, sql)
	return
}

// 学习模式
type Mode struct {
	ModeId   int
	ModeName string
}

func (m *Mode) get(modeId int) (res Mode) {
	res.ModeId = modeId
	switch modeId {
	case 1:
		res.ModeName = "背诵"
	case 2:
		res.ModeName = "默写"
	case 3:
		res.ModeName = "认识"
	case 4:
		res.ModeName = "听写"
	}
	return
}

func (m *Mode) list() (res []Mode) {
	res = append(res, m.get(1))
	res = append(res, m.get(2))
	res = append(res, m.get(3))
	res = append(res, m.get(4))
	return
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

func (l *Learn) insert() {
	sql := `INSERT INTO learn(mode_id, kind_id, deck_id, learn_name) VALUES(?, ?, 0, ?)`
	db.Exec(sql, l.ModeId, l.KindId, l.LearnName)
}

func (l *Learn) update() {
	sql := `UPDATE learn SET mode_id=?, kind_id=?, deck_id=?, learn_name=? WHERE learn_id=?`
	db.Exec(sql, l.ModeId, l.KindId, l.DeckId, l.LearnName, l.LearnId)
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
