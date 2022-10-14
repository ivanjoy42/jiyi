package main

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	card     Card
	deck     Deck
	cardDeck CardDeck
	kind     Kind
	learn    Learn
)

func main() {
	r := gin.Default()
	r.StaticFile("/favicon.ico", "favicon.ico")
	r.Static("/static", "static")
	r.LoadHTMLGlob("tpl/*")
	r.Use(setVer)
	r.GET("/", indexRoute)
	r.GET("/home", homeRoute)
	r.GET("/card", cardRoute)
	r.GET("/setting", settingRoute)
	r.GET("/user", userRoute)
	r.GET("/learn", learnRoute)
	cardGroup(r.Group("card"))
	deckGroup(r.Group("deck"))
	r.Run(":8080")
}

func indexRoute(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func homeRoute(c *gin.Context) {
	c.HTML(200, "home.html", gin.H{
		"Learn": learn.list(),
	})
}

func cardRoute(c *gin.Context) {
	c.HTML(200, "card.html", gin.H{})
}

func settingRoute(c *gin.Context) {
	c.HTML(200, "setting.html", gin.H{})
}

func userRoute(c *gin.Context) {
	c.HTML(200, "user.html", gin.H{})
}

func learnRoute(c *gin.Context) {
	learnId, _ := strconv.Atoi(c.Query("learnId"))
	c.HTML(200, "learn.html", gin.H{
		"Learn": learn.get(learnId),
	})
}

func setVer(c *gin.Context) {
	f1, _ := os.Stat("static/index.css")
	ts1 := f1.ModTime().Unix()

	f2, _ := os.Stat("static/index.js")
	ts2 := f2.ModTime().Unix()

	ver := strconv.FormatInt(ts1, 10)
	if ts2 > ts1 {
		ver = strconv.FormatInt(ts2, 10)
	}

	c.SetCookie("ver", ver, 86400*30, "", "", false, false)
}

// 卡片操作
//
// list列表页面；
// create新建页面，modify修改页面，remove删除页面；
// insert插入记录，update更新记录，delete删除记录。
func cardGroup(r *gin.RouterGroup) {
	r.GET("list", func(c *gin.Context) {
		kindId, _ := strconv.Atoi(c.Query("kindId"))
		c.HTML(200, "cardList.html", gin.H{
			"Card": card.list(kindId),
			"Kind": kind.get(kindId),
		})
	})

	r.GET("create", func(c *gin.Context) {
		kindId, _ := strconv.Atoi(c.Query("kindId"))
		c.HTML(200, "cardCreate.html", gin.H{
			"KindId": kindId,
		})
	})

	r.GET("modify", func(c *gin.Context) {
		cardId, _ := strconv.Atoi(c.Query("cardId"))
		c.HTML(200, "cardModify.html", gin.H{
			"Card": card.get(cardId),
		})
	})

	r.GET("remove", func(c *gin.Context) {
		cardId, _ := strconv.Atoi(c.Query("cardId"))
		c.HTML(200, "cardRemove.html", gin.H{
			"Card": card.get(cardId),
		})
	})

	r.GET("search", func(c *gin.Context) {
		kindId, _ := strconv.Atoi(c.Query("kindId"))
		query := c.Query("query")
		c.HTML(200, "cardSearch.html", gin.H{
			"Card":   card.search(kindId, query),
			"KindId": kindId,
			"Query":  query,
		})
	})

	r.POST("insert", func(c *gin.Context) {
		card.KindId, _ = strconv.Atoi(c.PostForm("kindId"))
		card.Front = c.PostForm("front")
		card.Back = c.PostForm("back")
		card.Helper = c.PostForm("helper")
		card.Pinyin = c.PostForm("pinyin")
		card.insert()
	})

	r.POST("update", func(c *gin.Context) {
		card.CardId, _ = strconv.Atoi(c.PostForm("cardId"))
		card.Front = c.PostForm("front")
		card.Back = c.PostForm("back")
		card.Helper = c.PostForm("helper")
		card.Pinyin = c.PostForm("pinyin")
		card.update()
	})

	r.POST("delete", func(c *gin.Context) {
		cardId, _ := strconv.Atoi(c.PostForm("cardId"))
		card.deleteTx(cardId)
	})
}

// 卡组操作
//
// list列表页面；
// create新建页面，modify修改页面，remove删除页面；
// insert插入记录，update更新记录，delete删除记录。
func deckGroup(r *gin.RouterGroup) {
	r.GET("list", func(c *gin.Context) {
		kindId, _ := strconv.Atoi(c.Query("kindId"))
		c.HTML(200, "deckList.html", gin.H{
			"Deck": deck.list(kindId),
			"Kind": kind.get(kindId),
		})
	})

	r.GET("create", func(c *gin.Context) {
		kindId, _ := strconv.Atoi(c.Query("kindId"))
		c.HTML(200, "deckCreate.html", gin.H{
			"KindId": kindId,
		})
	})

	r.GET("modify", func(c *gin.Context) {
		deckId, _ := strconv.Atoi(c.Query("deckId"))
		c.HTML(200, "deckModify.html", gin.H{
			"Deck":   deck.get(deckId),
			"Fronts": deck.getFronts(deckId),
		})
	})

	r.GET("remove", func(c *gin.Context) {
		deckId, _ := strconv.Atoi(c.Query("deckId"))
		c.HTML(200, "deckRemove.html", gin.H{
			"Deck":   deck.get(deckId),
			"Fronts": deck.getFronts(deckId),
		})
	})

	r.POST("insert", func(c *gin.Context) {
		deck.KindId, _ = strconv.Atoi(c.PostForm("kindId"))
		deck.DeckName = c.PostForm("deckName")
		fronts := c.PostForm("fronts")
		deck.insertTx(fronts)
	})

	r.POST("update", func(c *gin.Context) {
		deck.DeckId, _ = strconv.Atoi(c.PostForm("deckId"))
		deck.KindId, _ = strconv.Atoi(c.PostForm("kindId"))
		deck.DeckName = c.PostForm("deckName")
		fronts := c.PostForm("fronts")
		deck.updateTx(fronts)
	})

	r.POST("delete", func(c *gin.Context) {
		deckId, _ := strconv.Atoi(c.PostForm("deckId"))
		deck.deleteTx(deckId)
	})
}
