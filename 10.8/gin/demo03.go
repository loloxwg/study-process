package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	//router.LoadHTMLGlob("C:\\Users\\zhangxuanwei\\Documents\\-DailyStudy\\10.8\\templates\\*")
	router.LoadHTMLFiles("templates/index.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8080")
}
