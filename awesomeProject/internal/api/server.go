package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ForecastType struct {
	Image   string
	Header  string
	Details string
	More    string
}

func StartServer() {
	log.Println("Server start up")
	data := [3]ForecastType{
		{
			"term.svg",
			"Прогноз температуры",
			"Точность ±50℃",
			"Наши термометры самые точные в мире!!!! Купи прогноз, не пожалеешь!",
		}, {
			"term.svg",
			"Прогноз давления",
			"Точность ±42кПа",
			"Наши манометры самые точные в мире!!!! Купи прогноз, не пожалеешь!",
		}, {
			"term.svg",
			"Прогноз влажности",
			"Точность ±66%",
			"Наши термометры самые точные в мире!!!! Купи прогноз, не пожалеешь!",
		},
	}
	log.Println(data)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.LoadHTMLGlob("templates/*")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"choose": true,
		})
	})

	r.GET("/more", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"choose": false,
		})
	})

	r.Static("/image", "resources")
	r.Static("/styles", "styles")

	r.Run()

	log.Println("Server down")
}
