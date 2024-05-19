package utils

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"
)

func CronAction() {
	url := "http://0.0.0.0:8000/send-emails"
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Post(url, "application/json", bytes.NewBuffer([]byte(``)))
	if err != nil {
		log.Println(fmt.Errorf("failed to call endpoint: %w", err))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(fmt.Errorf("unexpected status code: %d", resp.StatusCode))
		return
	}

	log.Println("Endpoint called successfully")
}
