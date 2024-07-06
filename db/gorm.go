package db

import (
	"fmt"
	"log"
	"os"
	"sujana-be-web-go/domain"
	"sujana-be-web-go/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormClientStruct struct {
	DB *gorm.DB
}

var GormClient GormClientStruct

func NewGormClient() {
	NewMySQLClient()
}

func NewMySQLClient() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	if GormClient.DB == nil {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}

		GormClient = GormClientStruct{
			DB: db,
		}
		env := os.Getenv("ENV")
		if env != "test" {
			RegisterTableToMigrate(db)
		}
	}
}

func RegisterTableToMigrate(db *gorm.DB) {
	e := db.AutoMigrate(
		&domain.Account{},
		&domain.User{},
		&domain.Order{},
		&domain.Report{},
	)
	if e != nil {
		log.Fatal(e)
	}
	SeederAccountTable()
}

func SeederAccountTable() {
	password := "admin"
	password, _ = utils.HashPassword(password)
	account := domain.Account{
		Username: "admin",
		Password: password,
		Role:     "admin",
	}
	GormClient.DB.Create(&account)
}
