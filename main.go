package main

import (
	"log"
	"net/http"
	"os"

	"github.com/deadmanssnitch/go-dmswebhooks"
)

type Config struct {
	Token    string
	Room     string
	Hostname string
}

func main() {
	cfg := &Config{
		Token:    os.Getenv("HIPCHAT_TOKEN"),
		Room:     os.Getenv("HIPCHAT_ROOM"),
		Hostname: os.Getenv("HIPCHAT_HOSTNAME"),
	}

	// Default to HipChat cloud
	if cfg.Hostname == "" {
		cfg.Hostname = "api.hipchat.com"
	}

	// HIPCHAT_TOKEN is required and is used for authentication. This example
	// requires a User token as the add-on tokens require oauth.
	if cfg.Token == "" {
		log.Fatalf("Missing HIPCHAT_TOKEN. Generate a new User API token.")
	}

	// HIPCHAT_ROOM is required an specifies which room to notify with alerts.
	if cfg.Room == "" {
		log.Fatalf("Missing HIPCHAT_ROOM. Set this to the Room ID to post to.")
	}

	http.Handle("/", dmswebhooks.NewHandler(
		func(alert *dmswebhooks.Alert) error {
			notice := newNotificiation(alert)

			return notifyHipchat(cfg, notice)
		},
	))

	var port = "4000"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	log.Printf("Started Listening on tcp://0.0.0.0:%v\n", port)
	http.ListenAndServe(":"+port, nil)
}
