package main

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	ver      string
	card     Card
	deck     Deck
	cardDeck CardDeck
	dir      Dir
	mode     Mode
	pract    Pract
)

func main() {
	r := gin.Default()
	r.StaticFile("favicon.ico", "static/favicon.ico")
	r.Static("static", "static")
	r.Static("app", "app")
	r.Static("static2", "static2")
	r.Static("app2", "app2")
	r.LoadHTMLGlob("tpl/**/*")
	r.Use(setVer)
	r.GET("/", indexRoute)
	cardGroup(r.Group("card"))
	deckGroup(r.Group("deck"))
	dirGroup(r.Group("dir"))
	practGroup(r.Group("pract"))
	practDeckGroup(r.Group("practDeck"))
	settingGroup(r.Group("setting"))
	userGroup(r.Group("user"))
	cardAPI(r.Group("api/card"))
	deckAPI(r.Group("api/deck"))
	dirAPI(r.Group("api/dir"))
	practAPI(r.Group("api/pract"))
	modeAPI(r.Group("api/mode"))
	r.Run(":8080")
}

func indexRoute(c *gin.Context) {
	c.HTML(200, "public/index.html", gin.H{
		"Ver": ver,
	})
}

func cardAPI(r *gin.RouterGroup) {
	r.GET("list", func(c *gin.Context) {
		dirId, _ := strconv.Atoi(c.Query("dirId"))
		deckId, _ := strconv.Atoi(c.Query("deckId"))
		c.JSON(200, gin.H{
			"Card": card.list(dirId, deckId),
			"Dir":  dir.get(dirId),
			"Deck": deck.get(deckId),
		})
	})

	r.GET("modify", func(c *gin.Context) {
		cardId, _ := strconv.Atoi(c.Query("cardId"))
		c.JSON(200, gin.H{
			"Card": card.get(cardId),
			"Dir":  dir.get(card.DirId),
		})
	})

	r.GET("search", func(c *gin.Context) {
		dirId, _ := strconv.Atoi(c.Query("dirId"))
		query := c.Query("query")
		c.JSON(200, gin.H{
			"Card":  card.search(dirId, query),
			"DirId": dirId,
			"Query": query,
		})
	})

	r.POST("insert", func(c *gin.Context) {
		data := struct {
			Card   *Card
			DeckId int
		}{&card, 0}
		c.BindJSON(&data)
		card.insertTx(data.DeckId)
	})

	r.POST("update", func(c *gin.Context) {
		c.BindJSON(&card)
		card.update()
	})

	r.POST("delete", func(c *gin.Context) {
		cardId, _ := strconv.Atoi(c.Query("cardId"))
		card.deleteTx(cardId)
	})
}

func deckAPI(r *gin.RouterGroup) {
	r.GET("list", func(c *gin.Context) {
		dirId, _ := strconv.Atoi(c.Query("dirId"))
		c.JSON(200, gin.H{
			"Deck": deck.list(dirId),
			"Dir":  dir.get(dirId),
		})
	})

	r.GET("modify", func(c *gin.Context) {
		deckId, _ := strconv.Atoi(c.Query("deckId"))
		c.JSON(200, gin.H{
			"Deck":   deck.get(deckId),
			"Fronts": deck.getFronts(deckId),
		})
	})

	r.POST("insert", func(c *gin.Context) {
		data := struct {
			Deck   *Deck
			Fronts string
		}{&deck, ""}
		c.BindJSON(&data)
		deck.insertTx(data.Fronts)
	})

	r.POST("update", func(c *gin.Context) {
		data := struct {
			Deck   *Deck
			Fronts string
		}{&deck, ""}
		c.BindJSON(&data)
		deck.updateTx(data.Fronts)
	})

	r.POST("delete", func(c *gin.Context) {
		deckId, _ := strconv.Atoi(c.Query("deckId"))
		deck.deleteTx(deckId)
	})
}

func dirAPI(r *gin.RouterGroup) {
	r.GET("list", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Dir": dir.list(),
		})
	})

	r.GET("modify", func(c *gin.Context) {
		dirId, _ := strconv.Atoi(c.Query("dirId"))
		c.JSON(200, gin.H{
			"Dir": dir.get(dirId),
		})
	})

	r.POST("insert", func(c *gin.Context) {
		c.BindJSON(&dir)
		dir.insert(dir.KindId)
	})

	r.POST("update", func(c *gin.Context) {
		c.BindJSON(&dir)
		dir.update()
	})

	r.POST("delete", func(c *gin.Context) {
		dirId, _ := strconv.Atoi(c.Query("dirId"))
		dir.delete(dirId)
	})
}

func practAPI(r *gin.RouterGroup) {
	r.GET("list", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Pract": pract.list(),
		})
	})

	r.POST("insert", func(c *gin.Context) {
		c.BindJSON(&pract)
		pract.insert()
	})
}

func modeAPI(r *gin.RouterGroup) {
	r.GET("list", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Mode": mode.list(),
		})
	})
}

// 卡片操作
//
// list列表页面；
// create新建页面，modify修改页面，remove删除页面；
// insert插入记录，update更新记录，delete删除记录。
func cardGroup(r *gin.RouterGroup) {
	r.GET("list", func(c *gin.Context) {
		dirId, _ := strconv.Atoi(c.Query("dirId"))
		deckId, _ := strconv.Atoi(c.Query("deckId"))
		c.HTML(200, "card/list.html", gin.H{
			"Card": card.list(dirId, deckId),
			"Dir":  dir.get(dirId),
			"Deck": deck.get(deckId),
		})
	})

	r.GET("create", func(c *gin.Context) {
		dirId, _ := strconv.Atoi(c.Query("dirId"))
		deckId, _ := strconv.Atoi(c.Query("deckId"))
		c.HTML(200, "card/create.html", gin.H{
			"DeckId": deckId,
			"Dir":    dir.get(dirId),
		})
	})

	r.GET("modify", func(c *gin.Context) {
		cardId, _ := strconv.Atoi(c.Query("cardId"))
		c.HTML(200, "card/modify.html", gin.H{
			"Card": card.get(cardId),
			"Dir":  dir.get(card.DirId),
		})
	})

	r.GET("remove", func(c *gin.Context) {
		cardId, _ := strconv.Atoi(c.Query("cardId"))
		c.HTML(200, "card/remove.html", gin.H{
			"Card": card.get(cardId),
			"Dir":  dir.get(card.DirId),
		})
	})

	r.GET("search", func(c *gin.Context) {
		dirId, _ := strconv.Atoi(c.Query("dirId"))
		query := c.Query("query")
		c.HTML(200, "card/search.html", gin.H{
			"Card":  card.search(dirId, query),
			"DirId": dirId,
			"Query": query,
		})
	})

	r.POST("insert", func(c *gin.Context) {
		card.DirId, _ = strconv.Atoi(c.PostForm("dirId"))
		card.Front = c.PostForm("front")
		card.Back = c.PostForm("back")
		card.Helper = c.PostForm("helper")
		card.Pinyin = c.PostForm("pinyin")
		deckId, _ := strconv.Atoi(c.PostForm("deckId"))
		card.insertTx(deckId)
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
		dirId, _ := strconv.Atoi(c.Query("dirId"))
		c.HTML(200, "deck/list.html", gin.H{
			"Deck": deck.list(dirId),
			"Dir":  dir.get(dirId),
		})
	})

	r.GET("create", func(c *gin.Context) {
		dirId, _ := strconv.Atoi(c.Query("dirId"))
		c.HTML(200, "deck/create.html", gin.H{
			"DirId": dirId,
		})
	})

	r.GET("modify", func(c *gin.Context) {
		deckId, _ := strconv.Atoi(c.Query("deckId"))
		c.HTML(200, "deck/modify.html", gin.H{
			"Deck":   deck.get(deckId),
			"Fronts": deck.getFronts(deckId),
		})
	})

	r.GET("remove", func(c *gin.Context) {
		deckId, _ := strconv.Atoi(c.Query("deckId"))
		c.HTML(200, "deck/remove.html", gin.H{
			"Deck":   deck.get(deckId),
			"Fronts": deck.getFronts(deckId),
		})
	})

	r.POST("insert", func(c *gin.Context) {
		deck.DirId, _ = strconv.Atoi(c.PostForm("dirId"))
		deck.DeckName = c.PostForm("deckName")
		fronts := c.PostForm("fronts")
		deck.insertTx(fronts)
	})

	r.POST("update", func(c *gin.Context) {
		deck.DeckId, _ = strconv.Atoi(c.PostForm("deckId"))
		deck.DirId, _ = strconv.Atoi(c.PostForm("dirId"))
		deck.DeckName = c.PostForm("deckName")
		fronts := c.PostForm("fronts")
		deck.updateTx(fronts)
	})

	r.POST("delete", func(c *gin.Context) {
		deckId, _ := strconv.Atoi(c.PostForm("deckId"))
		deck.deleteTx(deckId)
	})
}

// 文件夹
func dirGroup(r *gin.RouterGroup) {
	r.GET("list", func(c *gin.Context) {
		c.HTML(200, "dir/list.html", gin.H{
			"Dir": dir.list(),
		})
	})

	r.GET("create", func(c *gin.Context) {
		c.HTML(200, "dir/create.html", gin.H{})
	})

	r.GET("modify", func(c *gin.Context) {
		dirId, _ := strconv.Atoi(c.Query("dirId"))
		c.HTML(200, "dir/modify.html", gin.H{
			"Dir": dir.get(dirId),
		})
	})

	r.GET("remove", func(c *gin.Context) {
		dirId, _ := strconv.Atoi(c.Query("dirId"))
		c.HTML(200, "dir/remove.html", gin.H{
			"Dir": dir.get(dirId),
		})
	})

	r.POST("insert", func(c *gin.Context) {
		dir.DirName = c.PostForm("dirName")
		kindId, _ := strconv.Atoi(c.PostForm("kindId"))
		dir.insert(kindId)
	})

	r.POST("update", func(c *gin.Context) {
		dir.DirId, _ = strconv.Atoi(c.PostForm("dirId"))
		dir.DirName = c.PostForm("dirName")
		dir.update()
	})

	r.POST("delete", func(c *gin.Context) {
		dirId, _ := strconv.Atoi(c.PostForm("dirId"))
		dir.delete(dirId)
	})
}

// 练习操作
func practGroup(r *gin.RouterGroup) {
	r.GET("list", func(c *gin.Context) {
		c.HTML(200, "pract/list.html", gin.H{
			"Pract": pract.list(),
		})
	})

	r.GET("detail", func(c *gin.Context) {
		practId, _ := strconv.Atoi(c.Query("practId"))
		pract = pract.get(practId)
		deck = deck.get(pract.DeckId)
		c.HTML(200, "pract/detail.html", gin.H{
			"Pract": pract,
			"Deck":  deck,
		})
	})

	r.GET("create", func(c *gin.Context) {
		c.HTML(200, "pract/create.html", gin.H{
			"Dir":  dir.list(),
			"Mode": mode.list(),
		})
	})

	r.POST("insert", func(c *gin.Context) {
		pract.ModeId, _ = strconv.Atoi(c.PostForm("modeId"))
		pract.DirId, _ = strconv.Atoi(c.PostForm("dirId"))
		pract.PractName = c.PostForm("practName")
		pract.insert()
	})
}

// 练习卡组操作
func practDeckGroup(r *gin.RouterGroup) {
	r.GET("list", func(c *gin.Context) {
		practId, _ := strconv.Atoi(c.Query("practId"))
		dirId, _ := strconv.Atoi(c.Query("dirId"))
		c.HTML(200, "practDeck/list.html", gin.H{
			"Deck":    deck.list(dirId),
			"Dir":     dir.get(dirId),
			"PractId": practId,
		})
	})

	r.GET("detail", func(c *gin.Context) {
		practId, _ := strconv.Atoi(c.Query("practId"))
		deckId, _ := strconv.Atoi(c.Query("deckId"))
		c.HTML(200, "practDeck/detail.html", gin.H{
			"Deck":    deck.get(deckId),
			"Fronts":  deck.getFronts(deckId),
			"PractId": practId,
		})
	})

	r.POST("update", func(c *gin.Context) {
		practId, _ := strconv.Atoi(c.PostForm("practId"))
		pract.get(practId)
		pract.DeckId, _ = strconv.Atoi(c.PostForm("deckId"))
		pract.update()
	})
}

func settingGroup(r *gin.RouterGroup) {
	r.GET("index", func(c *gin.Context) {
		c.HTML(200, "setting/index.html", gin.H{})
	})
}

func userGroup(r *gin.RouterGroup) {
	r.GET("index", func(c *gin.Context) {
		c.HTML(200, "user/index.html", gin.H{})
	})
}

func setVer(c *gin.Context) {
	f1, _ := os.Stat("static/index.css")
	ts1 := f1.ModTime().Unix()

	f2, _ := os.Stat("static/index.js")
	ts2 := f2.ModTime().Unix()

	ver = strconv.FormatInt(ts1, 10)
	if ts2 > ts1 {
		ver = strconv.FormatInt(ts2, 10)
	}

	c.SetCookie("ver", ver, 86400*30, "", "", false, false)
}
