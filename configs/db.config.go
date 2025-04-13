package config

import (
	"backend_project/internal/middleware/exception"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=postgres user=developer dbname=backend_project port=5432 sslmode=disable"
	loggerDb := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Info,
			SlowThreshold:             time.Second,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:      loggerDb,
		PrepareStmt: true,
	})
	exception.PanicLog(err)

	DB = db
	instance, _ := DB.DB()
	instance.SetMaxOpenConns(20)
	instance.SetMaxIdleConns(10)
	instance.SetConnMaxLifetime(30000)
}

func DBGroupTransaction() *gorm.DB {
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	return tx
}

func DBGroupCommit(tx *gorm.DB) error {
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("DBGroupCommit error:%w", err)
	}
	return nil
}
