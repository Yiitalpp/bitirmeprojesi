package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB_USERNAME = "root"
var DB_PASSWORD = "projectDbPass"
var DB_NAME = "databasepr"
var DB_HOST = "project-db"
var DB_PORT = "3306"

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {

	if len(os.Getenv("PROJECT_USER")) > 0 {
		DB_USERNAME = os.Getenv("PROJECT_USER")
	}

	if len(os.Getenv("PROJECT_HOST")) > 0 {
		DB_HOST = os.Getenv("PROJECT_HOST")
	}

	if len(os.Getenv("PROJECT_PASS")) > 0 {
		DB_PASSWORD = os.Getenv("PROJECT_PASS")
	}

	if len(os.Getenv("PROJECT_PORT")) > 0 {
		DB_PORT = os.Getenv("PROJECT_PORT")
	}

	if len(os.Getenv("PROJECT_NAME")) > 0 {
		DB_NAME = os.Getenv("PROJECT_NAME")
	}

	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}

	return db
}
