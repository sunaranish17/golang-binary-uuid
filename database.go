package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
)

//DBConnection -> return db instance
func DBConnection() (*gorm.DB, error) {
	USER := "root"
	PASS := "admin123"
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "test"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), //io writer
		logger.Config{
			SlowThreshold: time.Second, //Slow SQL threshold
			LogLevel:      logger.Info, //Log Level
			Colorful:      true,        //Disable color
		},
	)

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)

	return gorm.Open(mysql.Open(url), &gorm.Config{Logger: newLogger})
}
