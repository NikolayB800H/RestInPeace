package api

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ForecastType struct {
	Image   string
	Header  string
	Details string
	More    string
	Id      string
	Unit    string
}

func StartServer() {
	log.Println("Server start up")

	data := [3]ForecastType{
		{
			"term.svg",
			"Прогноз температуры",
			"Точность ±50",
			"Наши термометры самые точные в мире!!!! Купи прогноз, не пожалеешь!",
			"1",
			"℃",
		}, {
			"gau.svg",
			"Прогноз давления",
			"Точность ±42",
			"Наши манометры самые точные в мире!!!! Купи прогноз, не пожалеешь!",
			"2",
			"мм рт. ст.",
		}, {
			"rain.svg",
			"Прогноз влажности",
			"Точность ±66",
			"Наши термометры самые точные в мире!!!! Купи прогноз, не пожалеешь!",
			"3",
			"%",
		},
	}

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/more/:id", func(c *gin.Context) {
		id_str := c.Param("id")
		id, err := strconv.Atoi(id_str)
		if err != nil || id < 1 || id > 3 {
			c.String(http.StatusNotFound, "Такого прогноза нет!")
			return
		}
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"choose": false,
			"id":     data[id-1],
		})
	})

	r.GET("/", func(c *gin.Context) {
		value := c.Query("index")
		log.Println("index", value)
		var forecasts []ForecastType
		for i := 0; i < len(data); i++ {
			if strings.Contains(strings.ToLower(data[i].Header), strings.ToLower(value)) {
				forecasts = append(forecasts, data[i])
			}
		}
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"choose": true,
			"data":   forecasts,
			"value":  value,
		})
	})

	r.Static("/image", "resources")
	r.Static("/styles", "styles")

	r.Run()

	log.Println("Server down")
}
