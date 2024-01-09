package app

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "awesomeProject/docs"
	"awesomeProject/internal/app/config"
	"awesomeProject/internal/app/dsn"
	"awesomeProject/internal/app/redis"
	"awesomeProject/internal/app/repository"
	"awesomeProject/internal/app/role"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Application struct {
	repo        *repository.Repository
	config      *config.Config
	minioClient *minio.Client
	redisClient *redis.Client
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ginBodyLogMiddleware(c *gin.Context) {
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	c.Next()
	statusCode := c.Writer.Status()
	if statusCode >= 400 {
		//ok this is an request with error, let's make a record for it
		// now print body (or log in your preferred way)
		fmt.Println("Response body: " + blw.body.String() + "$")
	}
}

func createOrUpdateInfluencer(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	println(string(body))

	c.Request.Body = io.NopCloser(bytes.NewReader(body))
}

func (app *Application) Run() {
	r := gin.Default()
	r.Use(ErrorHandler())
	r.Use(ginBodyLogMiddleware)
	r.Use(createOrUpdateInfluencer)

	api := r.Group("/api")
	{
		// Услуги (виды данных)
		d := api.Group("/data_types")
		{
			d.GET("", app.WithAuthCheck(role.NotAuthorized, role.Client, role.Moderator), app.GetAllDataTypes)                                  // Список с поиском
			d.GET("/:data_type_id", app.WithAuthCheck(role.NotAuthorized, role.Client, role.Moderator), app.GetDataType)                        // Одна услуга
			d.DELETE("/:data_type_id", app.WithAuthCheck(role.Moderator), app.DeleteDataType)                                                   // Удаление
			d.PUT("/:data_type_id", app.WithAuthCheck(role.Moderator), app.ChangeDataType)                                                      // Изменение
			d.POST("", app.WithAuthCheck(role.Moderator), app.AddDataType)                                                                      // Добавление
			d.POST("/:data_type_id/add_to_forecast_application", app.WithAuthCheck(role.Client, role.Moderator), app.AddToForecastApplications) // Добавление в заявление // Связь (связь заявок на предсказания и видов данных)
		}
		// Заявления (заявки на предсказания)
		f := api.Group("/forecast_applications")
		{
			f.GET("", app.WithAuthCheck(role.Client, role.Moderator), app.GetAllForecastApplications)                                       // Список (отфильтровать по дате формирования и статусу)
			f.GET("/:application_id", app.WithAuthCheck(role.Client, role.Moderator), app.GetForecastApplication)                           // Одно заявление
			f.PUT("/update", app.WithAuthCheck(role.Client, role.Moderator), app.UpdateForecastApplication)                                 // Изменение (изменение/добавление начальной даты)
			f.DELETE("", app.WithAuthCheck(role.Client, role.Moderator), app.DeleteForecastApplication)                                     // Удаление
			f.PUT("/user_confirm", app.WithAuthCheck(role.Client, role.Moderator), app.UserConfirm)                                         // Сформировать создателем
			f.PUT("/:application_id/moderator_confirm", app.WithAuthCheck(role.Moderator), app.ModeratorConfirm)                            // Завершить или отклонить модератором
			f.DELETE("/delete_data_type/:data_type_id", app.WithAuthCheck(role.Client, role.Moderator), app.DeleteFromForecastApplications) // Изменение (удаление услуг)
			f.PUT("/set_input/:data_type_id", app.WithAuthCheck(role.Client, role.Moderator), app.SetInput)                                 // Изменение входных данных
			f.PUT("/:application_id/calculate", app.Calculate)
		}
		// Пользователи (авторизация)
		u := api.Group("/user")
		{
			u.POST("/sign_up", app.Register)
			u.POST("/login", app.Login)
			u.GET("/logout", app.Logout)
		}
	}
	// !АХТУНГ! Никакого https!!!
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // http://0.0.0.0:8084/swagger/index.html ~/go/bin/swag init -g cmd/server/main.go
	r.Static("/image", "./resources")
	r.Static("/styles", "styles")
	r.Run(fmt.Sprintf("%s:%d", app.config.ServiceHost, app.config.ServicePort)) // 0.0.0.0:8084 по умолчанию
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

	app.minioClient, err = minio.New(app.config.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4("", "", ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	app.redisClient, err = redis.New(app.config.Redis)
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
