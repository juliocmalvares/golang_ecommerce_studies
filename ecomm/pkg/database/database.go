package database

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var once sync.Once
var instance *gorm.DB

func DeliverDatabaseConnection() (*gorm.DB, error) {
	once.Do(func() {
		instance, _ = connectDb()
	})

	return instance, nil
}

func connectDb() (*gorm.DB, error) {
	str_conn := ""
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,          // Disable color
		},
	)
	if viper.GetString("ENV") == "prod" {

	} else if viper.GetString("ENV") == "dev" {
		str_conn = fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", "localhost", "55433", "core", "tavern", "secure_password123")
		gormDB, err := gorm.Open(postgres.Open(str_conn), &gorm.Config{Logger: newLogger})
		if err != nil {
			fmt.Println(err)
			panic("Error to connect the database")
		}

		err = migrations(gormDB)
		if err != nil {
			fmt.Println(err)

		}
		DB, _ := gormDB.DB()
		DB.SetMaxIdleConns(100)
		DB.SetMaxOpenConns(1000)
		DB.SetConnMaxLifetime(time.Minute * 20)
		return gormDB, err
	}
	return nil, nil
}

func migrations(db *gorm.DB) error {
	return nil
}
