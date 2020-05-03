package models

import "github.com/jinzhu/gorm"

// Subscription table storing urls to be fired upon
// event
type Subscription struct {
	gorm.Model
	Webhook string
	Headers string
	Event   string
}
