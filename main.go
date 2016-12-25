package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/deadmanssnitch/go-dmswebhooks"
	"github.com/goji/httpauth"
)

type Config struct {
	Token    string
	Room     string
	Hostname string
	Password string
}

func main() {
	var err error

	cfg := &Config{
		Token:    os.Getenv("HIPCHAT_TOKEN"),
		Room:     os.Getenv("HIPCHAT_ROOM"),
		Hostname: os.Getenv("HIPCHAT_HOSTNAME"),
		Password: os.Getenv("DMS_PASSWORD"),
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

	handler := dmswebhooks.NewHandler(
		func(alert *dmswebhooks.Alert) error {
			notice := newNotificiation(alert)

			return notifyHipchat(cfg, notice)
		},
	)

	if cfg.Password != "" {
		handler = httpauth.SimpleBasicAuth(cfg.Password, "")(handler)
	}

	var port = "4000"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	log.Printf("Started Listening on tcp://0.0.0.0:%v\n", port)

	server := http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Printf("HTTP Server Error: %q\n", err.Error())
	}
}
