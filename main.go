package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("tpl/*")

	r.GET("/", index)

	r.GET("cardList", cardList)
	r.GET("cardDetail", cardDetail)
	r.GET("cardCreate", cardCreate)
	r.GET("cardUpdate", cardUpdate)
	r.GET("cardDelete", cardDelete)

	deck(r.Group("deck"))

	r.Run(":8080")
}

func index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

//todo：搜索卡片、分组浏览
func cardList(c *gin.Context) {
	card := selectCard()
	c.HTML(200, "cardList.html", gin.H{"card": card})
}

func cardDetail(c *gin.Context) {
	cardId := c.Query("cardId")
	card := getCard(cardId)
	c.HTML(200, "cardDetail.html", gin.H{"card": card})
}

func cardCreate(c *gin.Context) {
	front := c.Query("front")
	back := c.Query("back")

	//无数据，显示模板；有数据，写入数据库
	if front == "" && back == "" {
		c.HTML(200, "cardCreate.html", gin.H{"title": "添加卡片"})
	} else {
		insertCard(front, back)
		c.String(200, "添加卡片：%s %s\n写入数据库...", front, back)
	}
}

func cardUpdate(c *gin.Context) {
	cardId := c.Query("cardId")
	front := c.Query("front")
	back := c.Query("back")

	//无数据，显示模板；有数据，写入数据库
	if front == "" && back == "" {
		card := getCard(cardId)
		c.HTML(200, "cardUpdate.html", gin.H{
			"title": "修改卡片",
			"card":  card,
		})
	} else {
		updateCard(cardId, front, back)
		c.String(200, "修改卡片：%s %s\n写入数据库...", front, back)
	}
}

func cardDelete(c *gin.Context) {
	cardId := c.Query("cardId")
	delCard(cardId)
	c.String(200, "删除卡片：%s\n...", cardId)
}

func deck(r *gin.RouterGroup) {
	r.GET("list", func(c *gin.Context) {
		deck := selectDeck()
		c.HTML(200, "deckList.html", gin.H{"deck": deck})
	})

	r.GET("detail", func(c *gin.Context) {
		deckId := c.Query("deckId")
		deck := getDeck(deckId)
		card := selectCardByDeckId(deckId)
		c.HTML(200, "deckDetail.html", gin.H{
			"deck": deck,
			"card": card,
		})
	})

	r.GET("create", func(c *gin.Context) {
		deck := c.Query("deck")
		if deck == "" {
			c.HTML(200, "deckCreate.html", gin.H{"title": "添加卡组"})
		} else {
			insertDeck(deck)
			c.String(200, "添加卡组%s\n写入数据库...", deck)
		}
	})

	r.GET("update", func(c *gin.Context) {
		deckId := c.Query("deckId")
		deckName := c.Query("deckName")
		kind := c.Query("kind")
		cards := c.Query("cards")
		updateDeck(deckId, deckName)
		updateCardDeck(deckId, kind, cards)
		c.String(200, "update deck...")
	})

	r.GET("delete", func(c *gin.Context) {
		c.String(200, "delete deck...")
	})
}
