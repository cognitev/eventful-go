package models

import (
	"github.com/jinzhu/gorm"
)

// Event Table for storing events to subscribe to
type Event struct {
	gorm.Model
	EventID     string
	RetryPolicy string
}

// NotifySubscribers trigger all http subscribers
// tied to the event being set
func (e Event) NotifySubscribers(payload map[string]string) {
	db := Connect()
	defer db.Close()
	var subs []Subscription
	db.Find(&subs, "event = ?", e.EventID)
	for i := 0; i < len(subs); i++ {
		subs[i].Notify(payload, e.RetryPolicy)
	}
}
