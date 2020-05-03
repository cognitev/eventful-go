package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // for mysql
)

// Connect to mysql databse instance
func Connect() *gorm.DB {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	connectionString := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, name)
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		print(err)
		panic("[DB] failed to connect database")
	} else {
		return db
	}
}
