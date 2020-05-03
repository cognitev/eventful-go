package tasks

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// PerformNotify scheduled task to send http post request to subscriber with provided data
func PerformNotify(webhook string, event string, payload map[string]string, headers map[string]string) {
	celeryClient := intialize()
	celeryClient.Register("worker.notify", notify)
	_, err := celeryClient.Delay("worker.notify", webhook, event, payload, headers)
	if err != nil {
		panic(err)
	}
}

func notify(webhook string, event string, payload map[string]string, headers map[string]string) {
	client := &http.Client{}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln(err)
	}
	req, err := http.NewRequest("POST", webhook, bytes.NewBuffer(payloadBytes))
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
}
