package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("tpl/*")

	router.GET("/", index)
	router.GET("/cardCreate", cardCreate)
	router.GET("/cardList", cardList)
	router.GET("/cardDetail", cardDetail)
	router.GET("/cardDelete", cardDelete)
	router.GET("/cardUpdate", cardUpdate)

	router.Run()
}

func index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

//todo: restful模式
func cardCreate(c *gin.Context) {
	front := c.Query("front")
	back := c.Query("back")

	//无数据，显示模板；有数据，写入数据库
	if front == "" && back == "" {
		c.HTML(200, "cardCreate.html", gin.H{
			"title": "添加卡片",
		})
	} else {
		insertCard(front, back)
		c.String(200, "添加卡片：%s %s\n写入数据库...", front, back)
	}
}

func cardList(c *gin.Context) {
	card := selectCard()
	c.HTML(200, "cardList.html", gin.H{"card": card})
}

func cardDetail(c *gin.Context) {
	cardId := c.Query("cardId")
	card := getCard(cardId)
	c.HTML(200, "cardDetail.html", gin.H{"card": card})
}

func cardDelete(c *gin.Context) {
	cardId := c.Query("cardId")
	delCard(cardId)
	c.String(200, "删除卡片：%s\n...", cardId)
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
