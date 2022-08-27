package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db()
	router := gin.Default()
	router.LoadHTMLGlob("tpl/*")

	router.GET("/", index)
	router.GET("/cardCreate", cardCreate)

	router.Run()
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

//todo: restful模式
func cardCreate(c *gin.Context) {
	front := c.Query("front")
	back := c.Query("back")

	//无数据，显示模板；有数据，写入数据库
	if front == "" && back == "" {
		c.HTML(http.StatusOK, "cardCreate.html", gin.H{
			"title": "添加卡片",
		})
	} else {
		c.String(http.StatusOK, "添加卡片：%s %s\n写入数据库...", front, back)
	}

}
