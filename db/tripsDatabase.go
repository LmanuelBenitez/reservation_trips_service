package db

import (
	"database/sql"
	"template_soa/libs/logger"
	"template_soa/models"

	mysqlConfig "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TripsDatabase struct {
	Connection *gorm.DB
}

func NewTripsDatabase(config *models.Configuration) *TripsDatabase {
	return &TripsDatabase{
		Connection: connectDB(config),
	}
}

func connectDB(config *models.Configuration) *gorm.DB {

	dbConfig := mysqlConfig.Config{
		User:   config.Database.User,
		Passwd: config.Database.Password,
		Net:    "tcp",
		Addr:   config.Database.Server + ":" + config.Database.Port,
		DBName: config.Database.Name,
	}

	db, err := gorm.Open(mysql.Open(dbConfig.FormatDSN()), &gorm.Config{
		SkipDefaultTransaction: false,
	})

	if err != nil {
		logger.Fatal("TripsDatabase", "connectDB", err.Error())
	}

	sqlDB, err := sql.Open("mysql", dbConfig.FormatDSN())

	if err != nil {
		logger.Fatal("TripsDatabase", "connectDB", err.Error())
	}

	sqlDB.SetMaxOpenConns(10)

	return db
}
