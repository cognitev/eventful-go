package migration

import (
	"github.com/cognitev/eventful-go/models"
	_ "github.com/jinzhu/gorm/dialects/mysql" // for mysql
)

// Migrate function to run migrations for all models
func Migrate() {
	db := models.Connect()
	defer db.Close()
	db.AutoMigrate(&models.Event{}, &models.Subscription{})

}
