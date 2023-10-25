package ds

import (
	// "gorm.io/gorm"
	"time"
)

type Users struct {
	UserId      uint   `gorm:"primaryKey"`
	Login       string `gorm:"size:256;not null"`
	Password    string `gorm:"size:256;not null"`
	IsModerator bool   `gorm:"not null"`
}

type DataTypes struct {
	DataTypeId     uint    `gorm:"primaryKey;not null;autoIncrement:false"`
	ImagePath      string  `gorm:"size:256;not null"`
	DataTypeName   string  `gorm:"size:128;not null"`
	Precision      float64 `gorm:"not null"`
	Description    string  `gorm:"size:1024;not null"`
	Unit           string  `gorm:"size:32;not null"`
	DataTypeStatus string  `gorm:"size:50;not null"` // Replace with Enum
}

type ForecastApplications struct {
	ApplicationId             uint       `gorm:"primaryKey"`
	ApplicationStatus         string     `gorm:"size:50;not null"` // Replace with Enum
	ApplicationCreationDate   time.Time  `gorm:"not null;type:date"`
	ApplicationFormationDate  *time.Time `gorm:"type:date"`
	ApplicationCompletionDate *time.Time `gorm:"type:date"`
	CreatorId                 uint       `gorm:"not null"`
	ModeratorId               uint       `gorm:"not null"`
	/*InputStartDate            string     `gorm:"not null;type:date"`*/

	Moderator Users `gorm:"foreignKey:ModeratorId"`
	Creator   Users `gorm:"foreignKey:CreatorId"`
}

type ConnectorAppsTypes struct {
	DataTypeId    uint    `gorm:"primaryKey;not null;autoIncrement:false"`
	ApplicationId uint    `gorm:"primaryKey;not null;autoIncrement:false"`
	InputFirst    float64 `gorm:"not null"`
	InputSecond   float64 `gorm:"not null"`
	InputThird    float64 `gorm:"not null"`
	Output        float64

	DataType    *DataTypes            `gorm:"foreignKey:DataTypeId"`
	Application *ForecastApplications `gorm:"foreignKey:ApplicationId"`
}
