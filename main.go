package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("tpl/*")
	r.Static("/static", "static")

	r.GET("/", index)
	r.GET("/home", home)
	card(r.Group("card"))
	deck(r.Group("deck"))

	r.Run(":8080")
}

func index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func home(c *gin.Context) {
	c.HTML(200, "home.html", gin.H{})
}

// 卡片操作
//
// list列表页面；
// create新建页面，modify修改页面，remove删除页面；
// insert插入记录，update更新记录，delete删除记录。
//
// todo：搜索卡片、分组浏览
func card(r *gin.RouterGroup) {
	r.GET("list", func(c *gin.Context) {
		card := selectCard()
		c.HTML(200, "cardList.html", gin.H{"card": card})
	})

	r.GET("create", func(c *gin.Context) {
		c.HTML(200, "cardCreate.html", gin.H{})
	})

	r.GET("modify", func(c *gin.Context) {
		cardId := c.Query("cardId")
		card := getCard(cardId)
		c.HTML(200, "cardModify.html", gin.H{"card": card})
	})

	r.GET("remove", func(c *gin.Context) {
		cardId := c.Query("cardId")
		card := getCard(cardId)
		c.HTML(200, "cardRemove.html", gin.H{"card": card})
	})

	r.GET("search", func(c *gin.Context) {
		front := c.Query("front")
		card := searchCard(front)
		c.HTML(200, "cardSearch.html", gin.H{"card": card, "front": front})
	})

	r.POST("insert", func(c *gin.Context) {
		front := c.PostForm("front")
		back := c.PostForm("back")
		insertCard(front, back)
	})

	r.POST("update", func(c *gin.Context) {
		cardId := c.PostForm("cardId")
		front := c.PostForm("front")
		back := c.PostForm("back")
		updateCard(cardId, front, back)
	})

	r.POST("delete", func(c *gin.Context) {
		cardId := c.PostForm("cardId")
		deleteCard(cardId)
	})
}

// 卡组操作
//
// list列表页面；
// create新建页面，modify修改页面，remove删除页面；
// insert插入记录，update更新记录，delete删除记录。
func deck(r *gin.RouterGroup) {
	r.GET("list", func(c *gin.Context) {
		deck := selectDeck()
		c.HTML(200, "deckList.html", gin.H{"deck": deck})
	})

	r.GET("create", func(c *gin.Context) {
		c.HTML(200, "deckCreate.html", gin.H{})
	})

	r.GET("modify", func(c *gin.Context) {
		deckId := c.Query("deckId")
		deck := getDeck(deckId)
		card := selectCardByDeckId(deckId)
		c.HTML(200, "deckModify.html", gin.H{"deck": deck, "card": card})
	})

	r.GET("remove", func(c *gin.Context) {
		deckId := c.Query("deckId")
		deck := getDeck(deckId)
		card := selectCardByDeckId(deckId)
		c.HTML(200, "deckRemove.html", gin.H{"deck": deck, "card": card})
	})

	r.POST("insert", func(c *gin.Context) {
		deckName := c.PostForm("deckName")
		insertDeck(deckName)
	})

	r.POST("update", func(c *gin.Context) {
		deckId := c.PostForm("deckId")
		deckName := c.PostForm("deckName")
		kind := c.PostForm("kind")
		cards := c.PostForm("cards")
		updateDeck(deckId, deckName)
		updateCardDeck(deckId, kind, cards)
	})

	r.POST("delete", func(c *gin.Context) {
		deckId := c.PostForm("deckId")
		deleteDeck(deckId)
	})
}
