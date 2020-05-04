package main

import (
	"testing"

	"github.com/cognitev/eventful-go/tasks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestEventDispatch(t *testing.T) {

}

func TestNotify(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("POST", "http://testnotify.com",
		httpmock.NewStringResponder(200, `ok`))

	url := "http://testnotify.com"
	event := "test_event"
	headers := map[string]string{"content-type": "json"}
	payload := make(map[string]interface{})
	payload["hello"] = "world"
	tasks.Notify(url, event, payload, headers)
	assert.Equal(t, httpmock.GetTotalCallCount(), 1, "Endpoint must be called one time.")
}

func main() {

}
