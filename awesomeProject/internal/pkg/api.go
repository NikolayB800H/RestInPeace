package app

import (
	"awesomeProject/internal/app/ds"
	"awesomeProject/internal/app/role"
	"awesomeProject/internal/schemes"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetAllDataTypes godoc
// @Summary      Запросить все виды данных прогнозов и черновик заявки на прогноз
// @Description  Список видов данных включает только те, что со статусом "доступен"
// @Tags         Виды данных
// @Produce      json
// @Success      200 {object} schemes.GetAllDataTypesResponse
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
	var draftForecastApplications *ds.ForecastApplications = nil
	if userId, exists := c.Get("userId"); exists {
		draftForecastApplications, err = app.repo.GetDraftForecastApplication(userId.(string))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
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

// GetDataType godoc
// @Summary      Запросить один вид данных прогнозов
// @Description  Возвращает более подробную информацию об одном виде данных
// @Tags         Виды данных
// @Param        data_type_id path string true "уникальный идентификатор вида данных"
// @Produce      json
// @Success      200 {object} ds.DataTypes
// @Router       /api/data_types/{data_type_id} [get]
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

// DeleteDataType godoc
// @Summary      Запросить удаление вида данных прогнозов
// @Description  Удаляет один вид данных по его data_type_id
// @Tags         Виды данных
// @Param        data_type_id path string true "уникальный идентификатор вида данных"
// @Produce      json
// @Success      200
// @Router       /api/data_types/{data_type_id} [delete]
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
	if dataType.ImagePath != nil {
		if err := app.deleteImage(c, dataType.DataTypeId); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	dataType.ImagePath = nil
	dataType.DataTypeStatus = ds.DELETED_TYPE
	if err := app.repo.SaveDataType(dataType); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

// AddDataType godoc
// @Summary      Запросить добавление вида данных прогнозов
// @Description  Добавляет один вид данных с заданными полями
// @Tags         Виды данных
// @Accept       mpfd
// @Param        image_path       formData file   false "Изображение вида данных"
// @Param        data_type_name   formData string true  "Название вида данных"
// @Param        precision        formData number true  "Погрешность предсказания вида данных"
// @Param        description      formData string true  "Описание вида данных"
// @Param        unit             formData string true  "Единица измерения вида данных"
// @Param        data_type_status formData string true  "Статус вида данных"
// @Produce      json
// @Success      200
// @Router       /api/data_types [post]
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

// ChangeDataType godoc
// @Summary      Запросить изменение вида данных прогнозов
// @Description  Изменяет один вид данных
// @Tags         Виды данных
// @Accept       mpfd
// @Param        data_type_id     path     string true  "уникальный идентификатор вида данных"
// @Param        image_path       formData file   false "Изображение вида данных"
// @Param        data_type_name   formData string true  "Название вида данных"
// @Param        precision        formData number true  "Погрешность предсказания вида данных"
// @Param        description      formData string true  "Описание вида данных"
// @Param        unit             formData string true  "Единица измерения вида данных"
// @Param        data_type_status formData string true  "Статус вида данных"
// @Produce      json
// @Success      200
// @Router       /api/data_types/{data_type_id} [put]
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
		if dataType.ImagePath != nil {
			if err := app.deleteImage(c, dataType.DataTypeId); err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		}
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

// AddToForecastApplications godoc
// @Summary      Запросить добавление вида данных в заявку на прогноз
// @Description  Добавляет данный вид данных в черновик заявки
// @Tags         Виды данных
// @Param        data_type_id path string true "уникальный идентификатор вида данных"
// @Produce      json
// @Success      200 {object} schemes.AllDataTypesResponse
// @Router       /api/data_types/{data_type_id} [post]
func (app *Application) AddToForecastApplications(c *gin.Context) {
	var request schemes.AddToForecastApplicationsRequest
	if err := c.ShouldBindUri(&request.URI); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	/*if err := c.ShouldBind(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}*/
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
	userId := getUserId(c)
	application, err = app.repo.GetDraftForecastApplication(userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if application == nil {
		application, err = app.repo.CreateDraftForecastApplication(userId)
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

// GetAllForecastApplications godoc
// @Summary      Запросить все заявки на прогнозы
// @Description  Возвращает все заявки с фильтрацией по статусу и дате формирования
// @Tags         Заявки на прогнозы
// @Param        status               query string false "статус заявки"
// @Param        formation_date_start query string false "начальная дата формирования"
// @Param        formation_date_end   query string false "конечная дата формирвания"
// @Produce      json
// @Success      200 {object} schemes.AllForecastApplicationssResponse
// @Router       /api/forecast_applications [get]
func (app *Application) GetAllForecastApplications(c *gin.Context) {
	var request schemes.GetAllForecastApplicationsRequest
	var err error
	if err := c.ShouldBindQuery(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId := getUserId(c)
	userRole := getUserRole(c)
	fmt.Println(userId, userRole)
	var applications []ds.ForecastApplications
	if userRole == role.Client {
		applications, err = app.repo.GetAllForecastApplications(&userId, request.FormationDateStart, request.FormationDateEnd, request.Status)
	} else {
		applications, err = app.repo.GetAllForecastApplications(nil, request.FormationDateStart, request.FormationDateEnd, request.Status)
	}
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

// GetForecastApplication godoc
// @Summary      Запросить одну заявку на прогноз
// @Description  Возвращает более подробную информацию о заявке
// @Tags         Заявки на прогнозы
// @Param        application_id path string true "уникальный идентификатор заявки"
// @Produce      json
// @Success      200 {object} schemes.ForecastApplicationsResponse
// @Router       /api/forecast_applications/{application_id} [get]
func (app *Application) GetForecastApplication(c *gin.Context) {
	var request schemes.ForecastApplicationRequest
	var err error
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	userId := getUserId(c)
	userRole := getUserRole(c)
	var application *ds.ForecastApplications
	if userRole == role.Moderator {
		application, err = app.repo.GetForecastApplicationById(request.ApplicationId, nil)
	} else {
		application, err = app.repo.GetForecastApplicationById(request.ApplicationId, &userId)
	}
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

// UpdateForecastApplication godoc
// @Summary      Запросить изменение черновика
// @Description  Изменяет дату начала входных измерений черновика и возвращает его
// @Tags         Заявки на прогнозы
// @Param        input_start_date query string false "дата начала входных измерений"
// @Produce      json
// @Success      200 {object} schemes.UpdateForecastApplicationsResponse
// @Router       /api/forecast_applications/ [put]
func (app *Application) UpdateForecastApplication(c *gin.Context) {
	var request schemes.UpdateForecastApplicationRequest
	var err error
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	userId := getUserId(c)
	var application *ds.ForecastApplications
	application, err = app.repo.GetDraftForecastApplication(userId)
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

// DeleteForecastApplication godoc
// @Summary      Удалить черновую заявку
// @Description  Удаляет черновую заявку пользователя
// @Tags         Заявки на прогнозы
// @Success      200
// @Router       /api/forecast_applications [delete]
func (app *Application) DeleteForecastApplication(c *gin.Context) {
	var err error
	var application *ds.ForecastApplications
	userId := getUserId(c)
	application, err = app.repo.GetDraftForecastApplication(userId)
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

// DeleteFromForecastApplications godoc
// @Summary      Запросить удаление вида данных из черновика заявки
// @Description  Удаляет один вид данных по его data_type_id
// @Tags         Заявки на прогнозы
// @Param        data_type_id path string true "уникальный идентификатор вида данных"
// @Produce      json
// @Success      200 {object} schemes.AllDataTypesResponse
// @Router       /api/forecast_applications/delete_data_type/{data_type_id} [delete]
func (app *Application) DeleteFromForecastApplications(c *gin.Context) {
	var request schemes.DeleteFromForecastApplicationsRequest
	var err error
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var application *ds.ForecastApplications
	userId := getUserId(c)
	application, err = app.repo.GetDraftForecastApplication(userId)
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

// UserConfirm godoc
// @Summary      Запросить формирование заявки
// @Description  Сформировать заявку пользователем
// @Tags         Заявки на прогнозы
// @Produce      json
// @Success      200 {object} schemes.UpdateForecastApplicationsResponse
// @Router       /api/forecast_applications/user_confirm [put]
func (app *Application) UserConfirm(c *gin.Context) {
	userId := getUserId(c)
	application, err := app.repo.GetDraftForecastApplication(userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if application == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("заявление не найдено"))
		return
	}
	if err := calculateRequest(application.ApplicationId); err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf(`сервис расчета прогноза не доступен: {%s}`, err))
		return
	}

	calculateStatus := ds.CalculateStarted
	application.CalculateStatus = &calculateStatus
	application.ApplicationStatus = ds.FORMED_APPLICATION
	now := time.Now()
	application.ApplicationFormationDate = &now

	if err := app.repo.SaveForecastApplication(application); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, schemes.UpdateForecastApplicationsResponse{ForecastApplications: schemes.ConvertForecastApplications(application)})
}

// ModeratorConfirm godoc
// @Summary      Подтвердить заявку
// @Description  Подтвердить или отменить заявку модератором
// @Tags         Заявки на прогнозы
// @Param        application_id path  string true  "уникальный идентификатор заявки"
// @Param        status         query string false "статус заявки"
// @Success      200 {object} schemes.UpdateForecastApplicationsResponse
// @Router       /api/forecast_applications/{application_id}/moderator_confirm [put]
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
	userId := getUserId(c)
	application, err := app.repo.GetForecastApplicationById(request.URI.ApplicationId, &userId)
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
	if request.Status == ds.COMPELTED_APPLICATION {
		now := time.Now()
		application.ApplicationCompletionDate = &now
	}
	moderator, err := app.repo.GetUserById(userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	application.ModeratorId = &userId
	application.Moderator = moderator

	if err := app.repo.SaveForecastApplication(application); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, schemes.UpdateForecastApplicationsResponse{ForecastApplications: schemes.ConvertForecastApplications(application)})
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
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId := getUserId(c)
	application, err := app.repo.GetDraftForecastApplication(userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if application == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("заявление не найдено"))
		return
	}
	if err := app.repo.SetInputConnectorAppsTypes(application.ApplicationId, request.URI.DataTypeId, request.InputFirst, request.InputSecond, request.InputThird); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (app *Application) Calculate(c *gin.Context) {
	var request schemes.CalculateReq
	if err := c.ShouldBindUri(&request.URI); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if request.Token != app.config.Token {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	application, err := app.repo.GetForecastApplicationById(request.URI.ApplicationId, nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if application == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("перевозка не найдена"))
		return
	}

	var calculateStatus string
	if *request.CalculateStatus {
		calculateStatus = ds.CalculateCompleted
	} else {
		calculateStatus = ds.CalculateFailed
	}
	application.CalculateStatus = &calculateStatus

	if err := app.repo.SaveForecastApplication(application); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
}
