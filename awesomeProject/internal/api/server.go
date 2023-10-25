/*
package api

import (

	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

)

	type ForecastType struct {
		Image        string
		ForecastType string
		Details      string
		More         string
		Id           string
		Unit         string
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
			value := c.Query("forecast")
			log.Println("forecast", value)
			var forecasts []ForecastType
			for i := 0; i < len(data); i++ {
				if strings.Contains(strings.ToLower(data[i].ForecastType), strings.ToLower(value)) {
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
*/
package api

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"awesomeProject/internal/app/models"
)

func StartServer() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/html/*")

	recipients := models.GetCardsInfo()

	r.GET("/recipients", func(c *gin.Context) {
		Name := c.Query("Name")
		filteredRecipients := filterRecipients(recipients, Name)

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			// "Recipients": recipients,
			"Recipients": filteredRecipients,
			"Name":       Name,
		})
	})

	r.GET("/recipients/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil || id < 0 || id >= len(recipients) {
			c.String(http.StatusNotFound, "Страница не найдена")
			return
		}
		recipient := recipients[id]

		c.HTML(http.StatusOK, "item.tmpl", gin.H{
			"Recipients": recipient,
		})
	})

	r.Static("/image", "./resources")
	//r.Static("/css", "./static/css")
	r.Static("/styles", "styles")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")

	r.Run()
	log.Println("Server down")
}

func filterRecipients(recipients []models.Recipients, filter string) []models.Recipients {
	if filter == "" {
		return recipients
	}
	var filtered []models.Recipients
	for _, recipient := range recipients {
		nameParts := strings.Fields(filter)
		matches := false
		for _, part := range nameParts {
			if contains(recipient.Name.First_name, part) || contains(recipient.Name.Second_name, part) {
				matches = true
				break
			}
		}
		if matches {
			filtered = append(filtered, recipient)
		}
	}

	return filtered
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr
}
