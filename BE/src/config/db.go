package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var db *gorm.DB

func DbConfig() {
	var err error
	dsn := "host=localhost user=postgress password=mbangg12 dbname=task-management port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed connect to database")
	}

	db.AutoMigrate()
}