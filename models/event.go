package models

import "github.com/jinzhu/gorm"

// Event Table for storing events to subscribe to
type Event struct {
	gorm.Model
	EventID     string
	RetryPolicy string
}
