package database

import (
	"Tugas2/models"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	dsn := "host=localhost user=postgres password=lupakatasandi dbname=pesanan port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connection to database :", err)
	}
	db.Debug().AutoMigrate(models.Order{}, models.Items{})

}

func GetDB() *gorm.DB {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("error connection to database :", err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
