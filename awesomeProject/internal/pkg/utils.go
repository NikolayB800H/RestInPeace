package app

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"mime/multipart"

	"awesomeProject/internal/app/role"
	"awesomeProject/internal/schemes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func (app *Application) uploadImage(c *gin.Context, image *multipart.FileHeader, UUID string) (*string, error) {
	src, err := image.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	extension := filepath.Ext(image.Filename)
	if extension != ".png" {
		return nil, fmt.Errorf("разрешены только png изображения")
	}
	imageName := UUID + extension

	_, err = app.minioClient.PutObject(c, app.config.Minio.BucketName, imageName, src, image.Size, minio.PutObjectOptions{
		ContentType: "image/png",
	})
	if err != nil {
		return nil, err
	}
	imageURL := fmt.Sprintf("http://%s/%s/%s", app.config.Minio.Endpoint, app.config.Minio.BucketName, imageName)
	return &imageURL, nil
}

func (app *Application) deleteImage(c *gin.Context, UUID string) error {
	imageName := UUID + ".png"
	//log.Println(imageName)
	err := app.minioClient.RemoveObject(c, app.config.Minio.BucketName, imageName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}

func getUserId(c *gin.Context) string {
	userId, _ := c.Get("userId")
	return userId.(string)
}

func getUserRole(c *gin.Context) role.Role {
	userRole, _ := c.Get("userRole")
	return userRole.(role.Role)
}

func generateHashString(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func (app *Application) calculateRequest(forecast_application_id string) error {
	url := "http://localhost:8000/api/calculate/"
	//var data schemes.CalculateAsyncReq
	data := &schemes.CalculateAsyncReq{
		ApplicationId: forecast_application_id,
		AllInputs:     []schemes.DataTypeInput{},
	}
	dataTypes, err := app.repo.GetConnectorAppsTypesExtended(forecast_application_id) //!!!
	if err != nil {
		return fmt.Errorf(`ошибка получения типов данных и входных значений: {%s}`, err)
	}
	for i, input := range dataTypes {
		log.Println(i, input)
		if input.InputFirst == nil || input.InputSecond == nil || input.InputThird == nil {
			return fmt.Errorf(`входное значение в заявке №{%d} не задано`, i)
		}
		data.AllInputs = append(data.AllInputs, schemes.DataTypeInput{
			DataTypeId:  input.DataTypeId,
			InputFirst:  *input.InputFirst,
			InputSecond: *input.InputSecond,
			InputThird:  *input.InputThird,
		})
	}
	payload, err := json.Marshal(data) //fmt.Sprintf(`{"application_id": "%s"}`, forecast_application_id)
	if err != nil {
		return fmt.Errorf(`формирование запроса сервису расчета прогноза провалено: {%s}`, err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(payload)) //bytes.NewBufferString(payload))
	if err != nil {
		return fmt.Errorf(`сервис расчета прогноза не доступен: {%s}`, err)
	}
	if resp.StatusCode >= 300 {
		return fmt.Errorf(`рассчёт провалился: {%s}`, resp.Status)
	}
	return nil
}
