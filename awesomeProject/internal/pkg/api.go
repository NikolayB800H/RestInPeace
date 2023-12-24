package app

import (
	"awesomeProject/internal/app/ds"
	"awesomeProject/internal/schemes"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"mime/multipart"

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

	_, err = app.minioClient.PutObject(c, app.config.BucketName, imageName, src, image.Size, minio.PutObjectOptions{
		ContentType: "image/jpeg",
	})
	if err != nil {
		return nil, err
	}
	imageURL := fmt.Sprintf("%s/%s/%s", app.config.MinioEndpoint, app.config.BucketName, imageName)
	return &imageURL, nil
}

func (app *Application) getCreator() string {
	return "5f58c307-a3f2-4b13-b888-c80ad08d5ed3"
}

func (app *Application) getModerator() *string {
	moderaorId := "796c70e1-5f27-4433-a415-95e7272effa5"
	return &moderaorId
}

// GetAllDataTypes godoc
// @Summary      Запросить все виды данных прогнозов и черновик заявки на прогноз
// @Description  Список видов данных включает только те, что со статусом "доступен"
// @Tags         Tests
// @Produce      json
// @Success      200  {object}  schemes.GetAllDataTypesResponse
// @Router       /api/data_types [get]
func (app *Application) GetAllDataTypes(c *gin.Context) {
	var request schemes.GetAllDataTypesRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	dataTypes, err := app.repo.GetDataTypeByName(request.DataTypeName)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	draftForecastApplications, err := app.repo.GetDraftForecastApplication(app.getCreator())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response := schemes.GetAllDataTypesResponse{DraftForecastApplications: nil, DataTypes: dataTypes}
	if draftForecastApplications != nil {
		response.DraftForecastApplications = &schemes.ForecastApplicationsShort{ApplicationId: draftForecastApplications.ApplicationId}
		dataTypes, err := app.repo.GetConnectorAppsTypes(draftForecastApplications.ApplicationId)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		response.DraftForecastApplications.DataTypeCount = len(dataTypes)
	}
	c.JSON(http.StatusOK, response)
}

func (app *Application) GetDataType(c *gin.Context) {
	var request schemes.DataTypeRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	dataType, err := app.repo.GetDataTypeByID(request.DataTypeId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if dataType == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("тип данных не найден"))
		return
	}
	c.JSON(http.StatusOK, dataType)
}

func (app *Application) DeleteDataType(c *gin.Context) {
	var request schemes.DataTypeRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	dataType, err := app.repo.GetDataTypeByID(request.DataTypeId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if dataType == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("тип данных не найден"))
		return
	}
	dataType.DataTypeStatus = ds.DELETED_TYPE
	if err := app.repo.SaveDataType(dataType); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (app *Application) AddDataType(c *gin.Context) {
	var request schemes.AddDataTypeRequest
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	dataType := ds.DataTypes(request.DataTypes)
	if err := app.repo.AddDataType(&dataType); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if request.ImagePath != nil {
		imagePath, err := app.uploadImage(c, request.ImagePath, dataType.DataTypeId)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		dataType.ImagePath = imagePath
	}
	if err := app.repo.SaveDataType(&dataType); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (app *Application) ChangeDataType(c *gin.Context) {
	var request schemes.ChangeDataTypeRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	dataType, err := app.repo.GetDataTypeByID(request.DataTypeId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if dataType == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("тип данных не найден"))
		return
	}
	if request.ImagePath != nil {
		imagePath, err := app.uploadImage(c, request.ImagePath, dataType.DataTypeId)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		dataType.ImagePath = imagePath
	}
	if request.DataTypeName != nil {
		dataType.DataTypeName = *request.DataTypeName
	}
	if request.Precision != nil {
		dataType.Precision = *request.Precision
	}
	if request.Description != nil {
		dataType.Description = *request.Description
	}
	if request.Unit != nil {
		dataType.Unit = *request.Unit
	}
	if request.DataTypeStatus != nil {
		dataType.DataTypeStatus = *request.DataTypeStatus
	}

	if err := app.repo.SaveDataType(dataType); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, dataType)
}

func (app *Application) AddToForecastApplications(c *gin.Context) {
	var request schemes.AddToForecastApplicationsRequest
	if err := c.ShouldBindUri(&request.URI); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var err error

	// Проверить существует ли тип данных
	dataType, err := app.repo.GetDataTypeByID(request.URI.DataTypeId) //!!!
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if dataType == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("тип данных не найден"))
		return
	}

	// Получить черновую заявку
	var application *ds.ForecastApplications
	application, err = app.repo.GetDraftForecastApplication(app.getCreator())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if application == nil {
		application, err = app.repo.CreateDraftForecastApplication(app.getCreator())
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	// Создать связь меджду заявкой и типом данных
	if err = app.repo.AddToConnectorAppsTypes(application.ApplicationId, request.URI.DataTypeId); err != nil { //!!!
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Вернуть список всех типов данных в заявке
	var dataTypes []ds.DataTypes
	dataTypes, err = app.repo.GetConnectorAppsTypes(application.ApplicationId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, schemes.AllDataTypesResponse{DataTypes: dataTypes})
}

func (app *Application) GetAllForecastApplications(c *gin.Context) {
	var request schemes.GetAllForecastApplicationsRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	applications, err := app.repo.GetAllForecastApplications(request.FormationDateStart, request.FormationDateEnd, request.Status)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	outputForecastApplications := make([]schemes.ForecastApplicationsOutput, len(applications))
	for i, application := range applications {
		outputForecastApplications[i] = schemes.ConvertForecastApplications(&application)
	}
	c.JSON(http.StatusOK, schemes.AllForecastApplicationssResponse{ForecastApplications: outputForecastApplications})
}

func (app *Application) GetForecastApplication(c *gin.Context) {
	var request schemes.ForecastApplicationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	application, err := app.repo.GetForecastApplicationById(request.ApplicationId, app.getCreator())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if application == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("заявление не найдено"))
		return
	}

	dataTypes, err := app.repo.GetConnectorAppsTypesExtended(request.ApplicationId) //!!!
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, schemes.ForecastApplicationsResponse{ForecastApplications: schemes.ConvertForecastApplications(application), DataTypes: dataTypes})
}

func (app *Application) UpdateForecastApplication(c *gin.Context) {
	var request schemes.UpdateForecastApplicationRequest
	if err := c.ShouldBindUri(&request.URI); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	application, err := app.repo.GetForecastApplicationById(request.URI.ApplicationId, app.getCreator())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if application == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("заявление не найдено"))
		return
	}
	application.InputStartDate = request.InputStartDate
	if app.repo.SaveForecastApplication(application); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, schemes.UpdateForecastApplicationsResponse{ForecastApplications: schemes.ConvertForecastApplications(application)})
}

func (app *Application) DeleteForecastApplication(c *gin.Context) {
	var request schemes.ForecastApplicationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	application, err := app.repo.GetForecastApplicationById(request.ApplicationId, app.getCreator())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if application == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("заявление не найдено"))
		return
	}
	application.ApplicationStatus = ds.DELETED_APPLICATION

	if err := app.repo.SaveForecastApplication(application); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
}

func (app *Application) DeleteFromForecastApplications(c *gin.Context) {
	var request schemes.DeleteFromForecastApplicationsRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	application, err := app.repo.GetForecastApplicationById(request.ApplicationId, app.getCreator())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if application == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("заявление не найдено"))
		return
	}
	if application.ApplicationStatus != ds.DRAFT_APPLICATION {
		c.AbortWithError(http.StatusMethodNotAllowed, fmt.Errorf("нельзя редактировать заявление со статусом: %s", application.ApplicationStatus))
		return
	}

	if err := app.repo.DeleteFromConnectorAppsTypes(request.ApplicationId, request.DataTypeId); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	dataTypes, err := app.repo.GetConnectorAppsTypes(request.ApplicationId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, schemes.AllDataTypesResponse{DataTypes: dataTypes})
}

func (app *Application) UserConfirm(c *gin.Context) {
	var request schemes.UserConfirmRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	application, err := app.repo.GetForecastApplicationById(request.ApplicationId, app.getCreator())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if application == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("заявление не найдено"))
		return
	}
	if application.ApplicationStatus != ds.DRAFT_APPLICATION {
		c.AbortWithError(http.StatusMethodNotAllowed, fmt.Errorf("нельзя сформировать заявление со статусом %s", application.ApplicationStatus))
		return
	}
	application.ApplicationStatus = ds.FORMED_APPLICATION
	now := time.Now()
	application.ApplicationFormationDate = &now

	if err := app.repo.SaveForecastApplication(application); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
}

func (app *Application) ModeratorConfirm(c *gin.Context) {
	var request schemes.ModeratorConfirmRequest
	if err := c.ShouldBindUri(&request.URI); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if request.Status != ds.COMPELTED_APPLICATION && request.Status != ds.REJECTED_APPLICATION {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("нельзя изменить статус на %s", request.Status))
		return
	}

	application, err := app.repo.GetForecastApplicationById(request.URI.ApplicationId, app.getCreator())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if application == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("заявление не найдено"))
		return
	}
	if application.ApplicationStatus != ds.FORMED_APPLICATION {
		c.AbortWithError(http.StatusMethodNotAllowed, fmt.Errorf("нельзя изменить статус с \"%s\" на \"%s\"", application.ApplicationStatus, request.Status))
		return
	}
	application.ApplicationStatus = request.Status
	application.ModeratorId = app.getModerator()
	if request.Status == ds.COMPELTED_APPLICATION {
		now := time.Now()
		application.ApplicationCompletionDate = &now
	}

	if err := app.repo.SaveForecastApplication(application); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
}

func (app *Application) SetOutput(c *gin.Context) {
	var request schemes.SetOutputRequest
	if err := c.ShouldBindUri(&request.URI); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := app.repo.SetOutputConnectorAppsTypes(request.URI.ApplicationId, request.URI.DataTypeId, request.Output); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (app *Application) SetInput(c *gin.Context) {
	var request schemes.SetInputRequest
	if err := c.ShouldBindUri(&request.URI); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := app.repo.SetInputConnectorAppsTypes(request.URI.ApplicationId, request.URI.DataTypeId, request.InputFirst, request.InputSecond, request.InputThird); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}
