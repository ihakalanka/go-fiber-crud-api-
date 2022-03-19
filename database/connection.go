package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"main.go/models"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	dsn := "root:iha075@tcp(127.0.0.1:3306)/test2?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Invalid database url")
	}
	sqldb, _ := connection.DB()

	err = sqldb.Ping()
	if err != nil {
		log.Fatal("Database connected")
	}
	fmt.Println("Database connection successful")
	return connection
}

func CloseDatabase(connection *gorm.DB) {
	sqldb, _ := connection.DB()
	sqldb.Close()
}

func Migrations() {
	Connect()
	CloseDatabase(Connect())
	Connect().AutoMigrate(&models.Book{})
}
