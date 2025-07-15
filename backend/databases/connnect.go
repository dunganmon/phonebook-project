package databases

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	dsn := "host=localhost user=postgres password=123456 dbname=phonebook port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	DB = db
	fmt.Println("Successfully connected to database")
}
