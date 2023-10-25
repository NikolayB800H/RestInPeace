package repository

import (
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"awesomeProject/internal/app/ds"
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
	dataType := &ds.DataTypes{}

	err := r.db.Where("data_type_id = ?", id).First(dataType).Error
	if err != nil {
		return nil, err
	}

	return dataType, nil
}

func (r *Repository) GetAllDataTypes() ([]ds.DataTypes, error) {
	var dataTypes []ds.DataTypes

	err := r.db.Find(&dataTypes).Error
	if err != nil {
		return nil, err
	}

	return dataTypes, nil
}

func (r *Repository) GetDataTypeByName(DataTypeName string) ([]ds.DataTypes, error) {
	var dataTypes []ds.DataTypes

	err := r.db.
		Where("LOWER(data_types.data_type_name) LIKE ?", "%"+strings.ToLower(DataTypeName)+"%").
		Find(&dataTypes).Error

	if err != nil {
		return nil, err
	}

	return dataTypes, nil
}

func (r *Repository) ApplicationSetStatus(id string) error {
	err := r.db.Exec("UPDATE data_types SET data_type_status = ? WHERE data_type_id = ?", "AMOGUS", id).Error
	if err != nil {
		return err
	}

	return nil
}
