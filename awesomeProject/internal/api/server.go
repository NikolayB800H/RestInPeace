package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	log.Println("Server start up")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.LoadHTMLGlob("../../templates/*")

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "base.layout.tmpl", gin.H{
			"title": "Main website",
		})
	})

	r.Static("/image", "../../resources")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}
