package ds

import (
	// "gorm.io/gorm"
	"awesomeProject/internal/app/role"
	"time"
)

type Users struct {
	UserId   string `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"-"`
	Login    string `gorm:"size:256;not null" json:"-"`
	Password string `gorm:"size:256;not null" json:"-"`
	Role     role.Role
}

const OK_TYPE string = "доступен"
const DELETED_TYPE string = "удалён"

type DataTypes struct {
	DataTypeId     string  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"data_type_id" binding:"-"`
	ImagePath      *string `gorm:"size:256" json:"image_path" binding:"-"`
	DataTypeName   string  `gorm:"size:128;not null" form:"data_type_name" json:"data_type_name" binding:"required,max=128"`
	Precision      float64 `gorm:"not null" form:"precision" json:"precision" binding:"required"`
	Description    string  `gorm:"size:1024;not null" form:"description" json:"description" binding:"required,max=1024"`
	Unit           string  `gorm:"size:32;not null" form:"unit" json:"unit" binding:"required,max=32"`
	DataTypeStatus string  `gorm:"size:50;not null" form:"data_type_status" json:"data_type_status" binding:"required,max=50"` // Replace with Enum
}

const DRAFT_APPLICATION string = "черновик"
const FORMED_APPLICATION string = "сформирован"
const COMPELTED_APPLICATION string = "завершён"
const REJECTED_APPLICATION string = "отклонён"
const DELETED_APPLICATION string = "удалён"

const CalculateCompleted string = "прогноз рассчитан"
const CalculateFailed string = "прогноз отменён"
const CalculateStarted string = "идёт рассчёт прогноза..."

type ForecastApplications struct {
	ApplicationId             string     `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ApplicationStatus         string     `gorm:"size:50;not null"` // Replace with Enum
	CalculateStatus           *string    `gorm:"size:50"`
	ApplicationCreationDate   time.Time  `gorm:"not null;type:timestamp"`
	ApplicationFormationDate  *time.Time `gorm:"type:timestamp"`
	ApplicationCompletionDate *time.Time `gorm:"type:timestamp"`
	CreatorId                 string     `gorm:"not null"`
	ModeratorId               *string    `json:"-"`
	InputStartDate            time.Time  `gorm:"not null;type:date"`

	Moderator *Users
	Creator   Users
}

type ConnectorAppsTypes struct {
	DataTypeId    string `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"data_type_id"`
	ApplicationId string `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"application_id"`
	InputFirst    *float64
	InputSecond   *float64
	InputThird    *float64
	Output        *float64

	DataType    *DataTypes            `gorm:"foreignKey:DataTypeId" json:"data_type"`
	Application *ForecastApplications `gorm:"foreignKey:ApplicationId" json:"application"`
}

type ConnectorAppsTypesDataTypes struct {
	DataTypeId     string  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"data_type_id" binding:"-"`
	ImagePath      *string `gorm:"size:256" json:"image_path" binding:"-"`
	DataTypeName   string  `gorm:"size:128;not null" form:"data_type_name" json:"data_type_name" binding:"required,max=128"`
	Precision      float64 `gorm:"not null" form:"precision" json:"precision" binding:"required"`
	Description    string  `gorm:"size:1024;not null" form:"description" json:"description" binding:"required,max=1024"`
	Unit           string  `gorm:"size:32;not null" form:"unit" json:"unit" binding:"required,max=32"`
	DataTypeStatus string  `gorm:"size:50;not null" form:"data_type_status" json:"data_type_status" binding:"required,max=50"` // Replace with Enum
	InputFirst     *float64
	InputSecond    *float64
	InputThird     *float64
	Output         *float64
} //!!!
