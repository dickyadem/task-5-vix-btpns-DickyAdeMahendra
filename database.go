package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func InitDatabase() {
	username := "root"     // username MySQL
	password := "password" // password MySQL
	dbName := "profile"    // nama database MySQL
	dbHost := "localhost"  // host MySQL

	dbUri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbName)

	conn, err := gorm.Open("mysql", dbUri)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	db = conn
}

func GetDB() *gorm.DB {
	return db
}
