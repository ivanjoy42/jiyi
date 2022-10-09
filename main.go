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
func card(r *gin.RouterGroup) {
	r.GET("list", func(c *gin.Context) {
		kindId := c.Query("kindId")
		c.HTML(200, "cardList.html", gin.H{
			"Card": selectCard(kindId),
			"Kind": getKind(kindId),
		})
	})

	r.GET("create", func(c *gin.Context) {
		kindId := c.Query("kindId")
		c.HTML(200, "cardCreate.html", gin.H{
			"KindId": kindId,
			"Kind":   getKind(kindId),
		})
	})

	r.GET("modify", func(c *gin.Context) {
		cardId := c.Query("cardId")
		c.HTML(200, "cardModify.html", gin.H{
			"Card": getCard(cardId),
		})
	})

	r.GET("remove", func(c *gin.Context) {
		cardId := c.Query("cardId")
		c.HTML(200, "cardRemove.html", gin.H{
			"Card": getCard(cardId),
		})
	})

	r.GET("search", func(c *gin.Context) {
		kindId := c.Query("kindId")
		front := c.Query("front")
		c.HTML(200, "cardSearch.html", gin.H{
			"Card":   searchCard(kindId, front),
			"Front":  front,
			"KindId": kindId,
		})
	})

	r.POST("insert", func(c *gin.Context) {
		kindId := c.PostForm("kindId")
		front := c.PostForm("front")
		back := c.PostForm("back")
		helper := c.PostForm("helper")
		pinyin := c.PostForm("pinyin")
		insertCard(kindId, front, back, helper, pinyin)
	})

	r.POST("update", func(c *gin.Context) {
		cardId := c.PostForm("cardId")
		front := c.PostForm("front")
		back := c.PostForm("back")
		helper := c.PostForm("helper")
		pinyin := c.PostForm("pinyin")
		updateCard(cardId, front, back, helper, pinyin)
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
		kindId := c.Query("kindId")
		c.HTML(200, "deckList.html", gin.H{
			"Deck": selectDeck(kindId),
			"Kind": getKind(kindId),
		})
	})

	r.GET("create", func(c *gin.Context) {
		kindId := c.Query("kindId")
		c.HTML(200, "deckCreate.html", gin.H{
			"KindId": kindId,
		})
	})

	r.GET("modify", func(c *gin.Context) {
		deckId := c.Query("deckId")
		c.HTML(200, "deckModify.html", gin.H{
			"Deck": getDeck(deckId),
			"Card": selectCardByDeckId(deckId),
		})
	})

	r.GET("remove", func(c *gin.Context) {
		deckId := c.Query("deckId")
		c.HTML(200, "deckRemove.html", gin.H{
			"Deck": getDeck(deckId),
			"Card": selectCardByDeckId(deckId),
		})
	})

	r.POST("insert", func(c *gin.Context) {
		kindId := c.PostForm("kindId")
		deckName := c.PostForm("deckName")
		cards := c.PostForm("cards")
		insertDeckTxn(deckName, kindId, cards)
	})

	r.POST("update", func(c *gin.Context) {
		deckId := c.PostForm("deckId")
		deckName := c.PostForm("deckName")
		kindId := c.PostForm("kindId")
		cards := c.PostForm("cards")
		updateDeckTxn(deckId, deckName, kindId, cards)
	})

	r.POST("delete", func(c *gin.Context) {
		deckId := c.PostForm("deckId")
		deleteDeck(deckId)
	})
}
