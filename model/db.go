package model

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	hostname := os.Getenv("POSTGRES_HOSTNAME")

	// set timezone
	timezone := "Asia/Tokyo"
	if value, ok := os.LookupEnv("TZ"); ok {
		timezone = value
	}
	dsn := "host=" + hostname + " user=" + user + " password=" + password + " dbname=" + dbname + " port=5432 sslmode=disable TimeZone=" + timezone
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	DB.AutoMigrate(&Post{})
}
