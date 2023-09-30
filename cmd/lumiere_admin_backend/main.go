package main

import (
	"log"
	"os"

	"github.com/p1xray/lumiere_admin_backend/internal/app"
	"github.com/p1xray/lumiere_admin_backend/internal/config"
)

// @title Lumiere admin API
// @version 1.0
// @description API sever for Lumiere admin

// @host localhost:8080
// @BasePath /
func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run() error {
	cfg := config.Read()
	app.Run(cfg)

	return nil
}
