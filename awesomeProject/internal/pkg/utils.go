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
	if extension != ".jpg" && extension != ".jpeg" {
		return nil, fmt.Errorf("разрешены только jpeg изображения")
	}
	imageName := UUID + extension

	_, err = app.minioClient.PutObject(c, app.config.Minio.BucketName, imageName, src, image.Size, minio.PutObjectOptions{
		ContentType: "image/jpeg",
	})
	if err != nil {
		return nil, err
	}
	imageURL := fmt.Sprintf("http://%s/%s/%s", app.config.Minio.Endpoint, app.config.Minio.BucketName, imageName)
	return &imageURL, nil
}

func (app *Application) deleteImage(c *gin.Context, UUID string) error {
	imageName := UUID + ".jpg"
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
		return err
	}
	for i, input := range dataTypes {
		log.Println(i, input)
		data.AllInputs = append(data.AllInputs, schemes.DataTypeInput{
			DataTypeId:  input.DataTypeId,
			InputFirst:  input.InputFirst,
			InputSecond: input.InputSecond,
			InputThird:  input.InputThird,
		})
	}
	payload, err := json.Marshal(data) //fmt.Sprintf(`{"application_id": "%s"}`, forecast_application_id)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(payload)) //bytes.NewBufferString(payload))
	if err != nil {
		return err
	}
	if resp.StatusCode >= 300 {
		return fmt.Errorf(`Calculate failed with status: {%s}`, resp.Status)
	}
	return nil
}
