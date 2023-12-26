package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"awesomeProject/internal/app/ds"
	"awesomeProject/internal/app/dsn"
)

func main() {
	_ = godotenv.Load()
	db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
	if err != nil {
		/*
			if strings.Contains(err.Error(), "does not exist") {

				createDatabaseCommand := fmt.Sprintf("CREATE DATABASE %s", os.Getenv("DB_NAME"))
				DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
				if err != nil {
					log.Println(result.Error)
					panic("failed to create database")
				}
				result := DB.Exec(createDatabaseCommand)
				if result.Error != nil {
					log.Println(result.Error)
					panic("failed to create database")
				}

				db, err = gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
				if err != nil {
		*/panic("failed to connect to database") /*
				}
			}
		*/
	}

	// Migrate the schema
	err = db.AutoMigrate(
		&ds.Users{},
		&ds.DataTypes{},
		&ds.ForecastApplications{},
		&ds.ConnectorAppsTypes{},
	)
	if err != nil {
		panic("cant migrate db")
	}
}
