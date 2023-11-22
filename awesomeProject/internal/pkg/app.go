package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"awesomeProject/internal/app/config"
	"awesomeProject/internal/app/dsn"
	"awesomeProject/internal/app/repository"
)

type Application struct {
	repo        *repository.Repository
	config      *config.Config
	minioClient *minio.Client
}

/*
	type GetDataTypesBack struct {
		DataTypes []ds.DataTypes
		Name      string
		Choose    bool
	}
*/
func (app *Application) Run() {
	r := gin.Default()
	r.Use(ErrorHandler())

	// Услуги (контейнеры)
	r.GET("/api/data_types", app.GetAllDataTypes)                                                      // Список с поиском
	r.GET("/api/data_types/:data_type_id", app.GetDataType)                                            // Одна услуга
	r.DELETE("/api/data_types/:data_type_id", app.DeleteDataType)                                      // Удаление
	r.PUT("/api/data_types/:data_type_id", app.ChangeDataType)                                         // Изменение
	r.POST("/api/data_types", app.AddDataType)                                                         // Добавление
	r.POST("/api/data_types/:data_type_id/add_to_forecast_application", app.AddToForecastApplications) // Добавление в заявление

	// Заявки (перевозки)
	r.GET("/api/forecast_applications", app.GetAllForecastApplications)                                                       // Список (отфильтровать по дате формирования и статусу)
	r.GET("/api/forecast_applications/:application_id", app.GetForecastApplication)                                           // Одно заявление
	r.PUT("/api/forecast_applications/:application_id/update", app.UpdateForecastApplication)                                 // Изменение (изменение/добавление начальной даты)
	r.DELETE("/api/forecast_applications/:application_id", app.DeleteForecastApplication)                                     // Удаление
	r.DELETE("/api/forecast_applications/:application_id/delete_data_type/:data_type_id", app.DeleteFromForecastApplications) // Изменение (удаление услуг)
	r.PUT("/api/forecast_applications/:application_id/user_confirm", app.UserConfirm)                                         // Сформировать создателем
	r.PUT("/api/forecast_applications/:application_id/moderator_confirm", app.ModeratorConfirm)                               // Завершить отклонить модератором

	r.Static("/image", "./resources")
	r.Static("/styles", "styles")
	r.Run(fmt.Sprintf("%s:%d", app.config.ServiceHost, app.config.ServicePort))
	log.Println("Server down")
}

func New() (*Application, error) {
	var err error
	loc, _ := time.LoadLocation("UTC")
	time.Local = loc
	app := Application{}
	app.config, err = config.NewConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	app.repo, err = repository.New(dsn.FromEnv())
	if err != nil {
		return nil, err
	}

	app.minioClient, err = minio.New(app.config.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4("", "", ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	return &app, nil
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			log.Println(err.Err)
		}
		lastError := c.Errors.Last()
		if lastError != nil {
			switch c.Writer.Status() {
			case http.StatusBadRequest:
				c.JSON(-1, gin.H{"error": "wrong request"})
			case http.StatusNotFound:
				c.JSON(-1, gin.H{"error": lastError.Error()})
			case http.StatusMethodNotAllowed:
				c.JSON(-1, gin.H{"error": lastError.Error()})
			default:
				c.Status(-1)
			}
		}
	}
}
