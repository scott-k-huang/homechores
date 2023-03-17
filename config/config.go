package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB = initializeDBConnection()

func initializeDBConnection() *gorm.DB {
	dsn := "host=localhost user=homechores password=h0m3ch0r35 dbname=homechores port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("DB Initialization Failure")
		return nil
	} else {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal("DB failed to initialize")
		} else {
			sqlDB.SetConnMaxLifetime(time.Hour)
			sqlDB.SetMaxIdleConns(10)
			sqlDB.SetMaxOpenConns(50)
		}
		return db
	}
}
