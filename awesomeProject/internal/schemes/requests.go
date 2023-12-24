package schemes

import (
	"awesomeProject/internal/app/ds"
	"mime/multipart"
	"time"
)

type DataTypeRequest struct {
	DataTypeId string `uri:"data_type_id" binding:"required,uuid"`
}

type GetAllDataTypesRequest struct {
	DataTypeName string `form:"data_type_name"`
}

type AddDataTypeRequest struct {
	ds.DataTypes
	ImagePath *multipart.FileHeader `form:"image_path" json:"image_path" binding:"required"`
}

type ChangeDataTypeRequest struct {
	DataTypeId     string                `uri:"data_type_id" binding:"required,uuid"`
	ImagePath      *multipart.FileHeader `form:"image_path" json:"image_path"`
	DataTypeName   *string               `form:"data_type_name" json:"data_type_name" binding:"omitempty,max=128"`
	Precision      *float64              `form:"precision" json:"precision"`
	Description    *string               `form:"description" json:"description" binding:"omitempty,max=1024"`
	Unit           *string               `form:"unit" json:"unit" binding:"omitempty,max=32"`
	DataTypeStatus *string               `form:"data_type_status" json:"data_type_status" binding:"omitempty,max=50"` // Replace with Enum
}

type AddToForecastApplicationsRequest struct {
	URI struct {
		DataTypeId string `uri:"data_type_id" binding:"required,uuid"`
	}
	//InputFirst  float64 `form:"input_first" json:"input_first" binding:"required"`
	//InputSecond float64 `form:"input_second" json:"input_second" binding:"required"`
	//InputThird  float64 `form:"input_third" json:"input_third" binding:"required"`
}

type GetAllForecastApplicationsRequest struct {
	FormationDateStart *time.Time `form:"formation_date_start" json:"formation_date_start" time_format:"2006-01-02 15:04:05"`
	FormationDateEnd   *time.Time `form:"formation_date_end" json:"formation_date_end" time_format:"2006-01-02 15:04:05"`
	Status             string     `form:"status" json:"status"`
}

type ForecastApplicationRequest struct {
	ApplicationId string `uri:"application_id" binding:"required,uuid"`
}

type UpdateForecastApplicationRequest struct {
	URI struct {
		ApplicationId string `uri:"application_id" binding:"required,uuid"`
	}
	InputStartDate time.Time `form:"input_start_date" json:"input_start_date" time_format:"2006-01-02"`
}

type DeleteFromForecastApplicationsRequest struct {
	ApplicationId string `uri:"application_id" binding:"required,uuid"`
	DataTypeId    string `uri:"data_type_id" binding:"required,uuid"`
}

type UserConfirmRequest struct {
	ApplicationId string `uri:"application_id" binding:"required,uuid"`
}

type ModeratorConfirmRequest struct {
	URI struct {
		ApplicationId string `uri:"application_id" binding:"required,uuid"`
	}
	Status string `form:"status" json:"status" binding:"required"`
}

type SetOutputRequest struct {
	URI struct {
		DataTypeId    string `uri:"data_type_id" binding:"required,uuid"`
		ApplicationId string `uri:"application_id" binding:"required,uuid"`
	}
	Output float64 `form:"output" json:"output" binding:"required"`
}

type SetInputRequest struct {
	URI struct {
		DataTypeId    string `uri:"data_type_id" binding:"required,uuid"`
		ApplicationId string `uri:"application_id" binding:"required,uuid"`
	}
	InputFirst  float64 `form:"input_first" json:"input_first" binding:"required"`
	InputSecond float64 `form:"input_second" json:"input_second" binding:"required"`
	InputThird  float64 `form:"input_third" json:"input_third" binding:"required"`
}
