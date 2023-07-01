package db

import (
	"firebond/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var dbInstance *gorm.DB

func GetDB() *gorm.DB {
	return dbInstance
}

func SetupPostgres() (*gorm.DB, error) {
	dbHost, dbPort, dbName, dbUser, dbPassword := os.Getenv("DEV_DATABASE_HOST"), os.Getenv("DEV_DATABASE_PORT"), os.Getenv("DEV_DATABASE_NAME"), os.Getenv("DEV_DATABASE_USERNAME"), os.Getenv("DEV_DATABASE_PASSWORD")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbName,
		dbPassword)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Println("Failed to connect database")
		return nil, err
	}

	errs := models.AutoMigrate(db)
	if errs != nil {
		log.Fatalf("Failed to create rates table: %v", err)
	}

	dbInstance = db
	log.Print("Connected to Database")

	return db, nil
}
