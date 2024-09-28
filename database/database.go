package database

import (
	"fmt"
	"github.com/orest-kostiuk/fiber-test/app/models"
	"github.com/orest-kostiuk/fiber-test/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {

	var err error
	dsn := config.Config("DB_URL")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error connecting to database")
	}

	DB = db

	fmt.Println("Connection Opened to Database")

	err = DB.AutoMigrate(&models.Post{})
	if err != nil {
		panic("Error migrating database")
	}
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("Error migrating database")
	}
	fmt.Println("Database Migrated")
}
