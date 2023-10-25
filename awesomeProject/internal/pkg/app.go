package pkg

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"awesomeProject/internal/app/config"
	"awesomeProject/internal/app/ds"
	"awesomeProject/internal/app/dsn"
	"awesomeProject/internal/app/repository"
)

type Application struct {
	repo   *repository.Repository
	config *config.Config
}

type GetDataTypesBack struct {
	DataTypes []ds.DataTypes
	Name      string
	Choose    bool
}

func (a *Application) Run() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		value := c.Query("forecast")
		log.Println("forecast", value)

		forecasts, err := a.repo.GetDataTypeByName(value)
		if err != nil {
			log.Println("Таких прогнозов нет!", err)
			c.Error(err)
			return
		}
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"choose": true,
			"data":   forecasts,
			"value":  value,
		})

	})

	r.GET("/more/:id", func(c *gin.Context) {
		id_str := c.Param("id")
		//id, err := strconv.Atoi(id_str)
		dataType, err := a.repo.GetDataTypeByID(id_str)
		if err != nil /* || id < 1 || id > 3*/ {
			c.String(http.StatusNotFound, "Такого прогноза нет!")
			return
		}
		if dataType.DataTypeStatus != "valid" {
			c.String(http.StatusNotFound, "Попытка доступа к удаленному!")
			return
		}
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"choose": false,
			"id":     *dataType,
		})
	})

	r.POST("/", func(c *gin.Context) {
		id := c.PostForm("delete")

		a.repo.ApplicationSetStatus(id)

		dataTypes, err := a.repo.GetDataTypeByName("")
		if err != nil {
			log.Println("Нет доступа!", err)
			c.Error(err)
			return
		}

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"choose": true,
			"data":   dataTypes,
			"value":  "",
		})
	})

	r.Static("/image", "./resources")
	r.Static("/styles", "styles")
	r.Run("localhost:9000")
	log.Println("Server down")
}

func New() (*Application, error) {
	var err error
	app := Application{}
	app.config, err = config.NewConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	app.repo, err = repository.New(dsn.FromEnv())
	if err != nil {
		return nil, err
	}

	return &app, nil
}
