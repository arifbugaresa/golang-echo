package config

import (
	"fmt"
	"golang-echo/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

/*
Inisialisasi database postgres
*/
func InitDB() {
	var err error

	dsn := "host=localhost user=postgres password=rahasia dbname=golang_echo port=5454 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println("Database Connected.")

	if err != nil {
		panic(err.Error())
	}

	initMigrate()
}

/*
Auto Migrate dengan gorm
*/
func initMigrate() {
	DB.AutoMigrate(&model.User{})
}
