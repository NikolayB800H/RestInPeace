package repository

import (
	"errors"
	"strings"
	"time"

	"awesomeProject/internal/app/ds"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	db *gorm.DB
}

func New(dsn string) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) GetDataTypeByID(id string) (*ds.DataTypes, error) { // ?
	dataType := &ds.DataTypes{DataTypeId: id}

	//err := r.db.Where("data_type_id = ?", id).First(dataType).Error
	err := r.db.First(dataType, "data_type_status = ?", ds.OK_TYPE).Error
	if err != nil {
		//return nil, err
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return dataType, nil
}

func (r *Repository) AddDataType(dataType *ds.DataTypes) error {
	err := r.db.Create(&dataType).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetDataTypeByName(dataTypeName string) ([]ds.DataTypes, error) {
	var dataTypes []ds.DataTypes

	err := r.db.
		//Where("LOWER(data_types.data_type_name) LIKE ? AND data_types.data_type_status = 'valid'", "%"+strings.ToLower(dataTypeName)+"%").
		Where("LOWER(data_type_name) LIKE ?", "%"+strings.ToLower(dataTypeName)+"%").
		Where("data_type_status = ?", ds.OK_TYPE).
		Order("data_type_id").
		Find(&dataTypes).Error

	if err != nil {
		return nil, err
	}

	return dataTypes, nil
}

func (r *Repository) SaveDataType(dataType *ds.DataTypes) error {
	err := r.db.Save(dataType).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) AddToConnectorAppsTypes(applicationId string, dataTypeId string) error {
	connector := ds.ConnectorAppsTypes{ApplicationId: applicationId, DataTypeId: dataTypeId /*, InputFirst: 0.0, InputSecond: 0.0, InputThird: 0.0*/}
	err := r.db.Create(&connector).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetAllForecastApplications(creatorId *string, formationDateStart, formationDateEnd *time.Time, status string) ([]ds.ForecastApplications, error) {
	var forecastApplications []ds.ForecastApplications
	query := r.db.
		Preload("Creator").
		Preload("Moderator").
		Where("LOWER(application_status) LIKE ?", "%"+strings.ToLower(status)+"%").
		Where("application_status != ? AND application_status != ?", ds.DELETED_APPLICATION, ds.DRAFT_APPLICATION).
		Order("application_creation_date")
	if creatorId != nil {
		query = query.Where("creator_id = ?", *creatorId)
	}
	if formationDateStart != nil && formationDateEnd != nil {
		query = query.Where("application_formation_date BETWEEN ? AND ?", *formationDateStart, *formationDateEnd)
	} else if formationDateStart != nil {
		query = query.Where("application_formation_date >= ?", *formationDateStart)
	} else if formationDateEnd != nil {
		query = query.Where("application_formation_date <= ?", *formationDateEnd)
	}

	if err := query.Find(&forecastApplications).Error; err != nil {
		return nil, err
	}
	return forecastApplications, nil
}

func (r *Repository) GetDraftForecastApplication(creatorId string) (*ds.ForecastApplications, error) {
	application := &ds.ForecastApplications{}
	err := r.db.First(application, ds.ForecastApplications{ApplicationStatus: ds.DRAFT_APPLICATION, CreatorId: creatorId}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return application, nil
}

func (r *Repository) CreateDraftForecastApplication(creatorId string) (*ds.ForecastApplications, error) {
	application := &ds.ForecastApplications{ApplicationCreationDate: time.Now(), CreatorId: creatorId, ApplicationStatus: ds.DRAFT_APPLICATION}
	err := r.db.Create(application).Error
	if err != nil {
		return nil, err
	}
	return application, nil
}

func (r *Repository) GetForecastApplicationById(forecastApplicationId string, creatorId *string) (*ds.ForecastApplications, error) {
	application := &ds.ForecastApplications{}
	query := r.db.Preload("Moderator").Preload("Creator").
		Where("application_status != ?", ds.DELETED_APPLICATION)
	if creatorId != nil {
		query = query.Where("creator_id = ?", creatorId)
	}
	err := query.First(application, ds.ForecastApplications{ApplicationId: forecastApplicationId}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return application, nil
}

func (r *Repository) GetConnectorAppsTypes(applicationId string) ([]ds.DataTypes, error) {
	var dataTypes []ds.DataTypes

	err := r.db.Table("connector_apps_types").
		Select("data_types.*").
		Joins("JOIN data_types ON connector_apps_types.data_type_id = data_types.data_type_id").
		Where(ds.ConnectorAppsTypes{ApplicationId: applicationId}).
		Scan(&dataTypes).Error

	if err != nil {
		return nil, err
	}
	return dataTypes, nil
}

func (r *Repository) GetConnectorAppsTypesExtended(applicationId string) ([]ds.ConnectorAppsTypesDataTypes, error) {
	var dataTypes []ds.ConnectorAppsTypesDataTypes

	err := r.db.Table("connector_apps_types").
		Select("*").
		Joins("JOIN data_types ON connector_apps_types.data_type_id = data_types.data_type_id").
		Where(ds.ConnectorAppsTypes{ApplicationId: applicationId}).
		Order("data_types.data_type_id").
		Scan(&dataTypes).Error

	if err != nil {
		return nil, err
	}
	return dataTypes, nil
}

func (r *Repository) SaveForecastApplication(application *ds.ForecastApplications) error {
	err := r.db.Save(application).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteFromConnectorAppsTypes(applicationId, dataTypeId string) error {
	err := r.db.Delete(&ds.ConnectorAppsTypes{ApplicationId: applicationId, DataTypeId: dataTypeId}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) SetOutputConnectorAppsTypes(applicationId string, dataTypeId string, output float64) error {
	err := r.db.Model(ds.ConnectorAppsTypes{}).Where("application_id = ? AND data_type_id = ?", applicationId, dataTypeId).Updates(ds.ConnectorAppsTypes{Output: &output}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) SetInputConnectorAppsTypes(applicationId string, dataTypeId string, inputFirst float64, inputSecond float64, inputThird float64) error {
	err := r.db.Model(ds.ConnectorAppsTypes{}).Where("application_id = ? AND data_type_id = ?", applicationId, dataTypeId).Updates(ds.ConnectorAppsTypes{InputFirst: &inputFirst, InputSecond: &inputSecond, InputThird: &inputThird}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) AddUser(user *ds.Users) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetUserByLogin(login string) (*ds.Users, error) {
	user := &ds.Users{}
	if err := r.db.Where("login = ?", login).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *Repository) GetUserById(user_id string) (*ds.Users, error) {
	user := &ds.Users{}
	if err := r.db.Where("user_id = ?", user_id).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *Repository) GetConnectorAppsTypesExtendedCounts(applicationId string) (int64, int64, error) {
	var total int64

	err := r.db.Table("connector_apps_types").
		Select("*").
		Joins("JOIN data_types ON connector_apps_types.data_type_id = data_types.data_type_id").
		Where(ds.ConnectorAppsTypes{ApplicationId: applicationId}).
		Count(&total).Error

	if err != nil {
		return 0, 0, err
	}

	var calculated int64

	err = r.db.Table("connector_apps_types").
		Select("*").
		Joins("JOIN data_types ON connector_apps_types.data_type_id = data_types.data_type_id").
		Where(ds.ConnectorAppsTypes{ApplicationId: applicationId}).
		Where("output IS NOT NULL").
		Count(&calculated).Error

	if err != nil {
		return 0, 0, err
	}
	return total, calculated, nil
}
