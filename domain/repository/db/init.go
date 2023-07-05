package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joshuaetim/akiraka3/domain/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DB() *gorm.DB {
	var dialect gorm.Dialector
	var dsn string

	switch os.Getenv("DB_DRIVER") {
	case "sqlite":
		dsn = os.Getenv("DATABASE_URL")
		if mem := os.Getenv("SQLITE_MEMORY"); mem != "" {
			dialect = sqlite.Open(mem)
		} else {
			dialect = sqlite.Open(dsn)
		}
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_DATABASE"))
		if os.Getenv("DB_URL") != "" {
			dsn = os.Getenv("DB_URL")
		}
		dialect = mysql.Open(dsn)
	default:
		log.Fatalf("invalid driver: %s", os.Getenv("DB_DRIVER"))
	}

	config := &gorm.Config{
		TranslateError: true,
	}
	if os.Getenv("ENV") == "dev" {
		config.Logger = logger.Default.LogMode(logger.Info)
	}
	db, err := gorm.Open(dialect, config)
	if err != nil {
		log.Fatalf("Error connecting to database (%v)(%v): %v", dialect.Name(), dsn, err)
		return nil
	}

	db.AutoMigrate(&model.User{}, &model.Photo{}, &model.Like{})
	return db
}
