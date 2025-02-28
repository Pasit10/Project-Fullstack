package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

var DataBase DB

func InitDatabaseConnection() {
	ConnectionMasterDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
	)

	db, err := gorm.Open(mysql.Open(ConnectionMasterDB), &gorm.Config{})
	if err != nil {
		log.Fatalf("[Database] Failed to connect to database: %s", err.Error())
	}

	DataBase.DB = db
}
