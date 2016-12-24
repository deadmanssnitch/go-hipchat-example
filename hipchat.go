package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	dwh "github.com/deadmanssnitch/go-dmswebhooks"
)

type Notification struct {
	From    string `json:"from"`
	Format  string `json:"message_format"`
	Message string `json:"message"`
	Color   string `json:"color"`
}

// new newNotificiation creates an appropriate notificiaton from a webhook
// alert.
func newNotificiation(alert *dwh.Alert) *Notification {
	notice := &Notification{
		From:   "Dead Man's Snitch",
		Format: "text",
	}

	snitch := alert.Data.Snitch

	// Set colors and message based on the type of the alert
	switch alert.Type {
	case dwh.TypeSnitchReporting:
		notice.Color = "green"
		notice.Message = fmt.Sprintf("🎉  %s is reporting", snitch.Name)

	case dwh.TypeSnitchErrored:
		notice.Color = "red"
		notice.Message = fmt.Sprintf("🚨  %s has errored", snitch.Name)

	case dwh.TypeSnitchMissing:
		notice.Color = "yello"
		notice.Message = fmt.Sprintf("❓  %s is missing", snitch.Name)
	}

	// TODO: Add tags
	// TODO: Add a link to the snitch
	// TODO: Add other stuff

	return notice
}

func notifyHipchat(cfg *Config, notice *Notification) error {
	var err error

	body := &bytes.Buffer{}
	encoder := json.NewEncoder(body)

	// Convert the Notification to JSON for sending over the wire
	if err = encoder.Encode(notice); err != nil {
		return err
	}

	// Create a custom http client so we can control the timeout.
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	uri := fmt.Sprintf("https://%s/v2/room/%s/notification", cfg.Hostname, cfg.Room)
	req, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return err
	}

	// HipChat uses a special "Bearer" type for authorization.
	req.Header.Set("Authorization", "Bearer "+cfg.Token)
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// HipChat should always send back a 204 No Content response but lets be
	// generous.
	if resp.StatusCode >= 300 {
		return fmt.Errorf("HipChat responded with %v", resp.StatusCode)
	}

	return nil
}
