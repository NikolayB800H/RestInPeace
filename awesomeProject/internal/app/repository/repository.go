package repository

import (
	"errors"
	"strings"
	"time"

	"awesomeProject/internal/app/ds"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(dsn string) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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

func (r *Repository) AddToConnectorAppsTypes(applicationId, dataTypeId string) error {
	connector := ds.ConnectorAppsTypes{ApplicationId: applicationId, DataTypeId: dataTypeId}
	err := r.db.Create(&connector).Error
	if err != nil {
		return err
	}
	return nil
}

//////////

func (r *Repository) GetAllForecastApplications(formationDateStart, formationDateEnd *time.Time, status string) ([]ds.ForecastApplications, error) {
	var forecastApplications []ds.ForecastApplications

	query := r.db.
		Preload("Customer").
		Preload("Moderator").
		Where("LOWER(application_status) LIKE ?", "%"+strings.ToLower(status)+"%").
		Where("status != ?", ds.DELETED_APPLICATION)
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

//!!!

func (r *Repository) GetDraftTransportation(customerId string) (*ds.Transportation, error) {
	transportation := &ds.Transportation{}
	err := r.db.First(transportation, ds.Transportation{Status: ds.DRAFT, CustomerId: customerId}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return transportation, nil
}

func (r *Repository) CreateDraftTransportation(customerId string) (*ds.Transportation, error) {
	transportation := &ds.Transportation{CreationDate: time.Now(), CustomerId: customerId, Status: ds.DRAFT}
	err := r.db.Create(transportation).Error
	if err != nil {
		return nil, err
	}
	return transportation, nil
}

func (r *Repository) GetTransportationById(transportationId, customerId string) (*ds.Transportation, error) {
	transportation := &ds.Transportation{}
	err := r.db.Preload("Moderator").Preload("Customer").
		Where("status != ?", ds.DELETED).
		First(transportation, ds.Transportation{UUID: transportationId, CustomerId: customerId}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return transportation, nil
}

func (r *Repository) GetTransportatioinComposition(transportationId string) ([]ds.Container, error) {
	var containers []ds.Container

	err := r.db.Table("transportation_compositions").
		Select("containers.*").
		Joins("JOIN containers ON transportation_compositions.container_id = containers.uuid").
		Where(ds.TransportationComposition{TransportationId: transportationId}).
		Scan(&containers).Error

	if err != nil {
		return nil, err
	}
	return containers, nil
}

func (r *Repository) SaveTransportation(transportation *ds.Transportation) error {
	err := r.db.Save(transportation).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteFromTransportation(transportationId, ContainerId string) error {
	err := r.db.Delete(&ds.TransportationComposition{TransportationId: transportationId, ContainerId: ContainerId}).Error
	if err != nil {
		return err
	}
	return nil
}
