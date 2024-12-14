package main

import (
	"log"

	"github.com/dariuszdroba/go-from-template/config"
	"github.com/dariuszdroba/go-from-template/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
