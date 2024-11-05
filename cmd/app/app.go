package main

import (
	"log"
	"tournament_scoring/config"
	"tournament_scoring/internal/app"
	"tournament_scoring/pkg/logger"
)

func main() {
	cfg, err := config.New("./config/config.yml")
	if err != nil {
		log.Fatalf("new config: %s", err)
	}
	l := logger.New(cfg.Log.Level)
	l.Debug("debug messages are enabled")
	app.Migrate(cfg, l)
	app.Run(cfg, l)
}
