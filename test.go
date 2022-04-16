package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/abc", abc)
	router.GET("/", index)
	router.Run()
}

func abc(c *gin.Context) {
	c.String(200, "hello world")
}

func index(c *gin.Context) {
	c.String(200, "你好")
}
