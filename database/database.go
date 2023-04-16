package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/Hamid-Afify9/Smart-API-With-Go-Docker-and-Postgres/models"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Africa/Cairo",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		//panic(err)
		log.Fatal("Error connecting to database \n ", err)
		os.Exit(2)
	}
	log.Println("Database connected successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Database migration started")
	db.AutoMigrate(&models.Fact{})

	DB = Dbinstance {
		Db : db,
	
	}
}

