package models

import (
	"encoding/json"

	"github.com/cognitev/eventful-go/tasks"

	"github.com/jinzhu/gorm"
)

// Subscription table storing urls to be fired upon
// event
type Subscription struct {
	gorm.Model
	Webhook string
	Headers string
	Event   string
}

// Notify : Trigger Subscriber by calling its webhook
func (s Subscription) Notify(payload map[string]string, retryPolicy string) {
	var headers map[string]string
	if err := json.Unmarshal([]byte(s.Headers), &headers); err != nil {
		panic(err)
	}
	tasks.PerformNotify(s.Webhook, s.Event, payload, headers)
}
