package main

import (
	"log"
	"os"

	"github.com/p1xray/lumiere_admin_backend/internal/config"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run() error {
	_ = config.Read()

	return nil
}
