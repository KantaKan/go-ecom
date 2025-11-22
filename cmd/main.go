package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}

	api := application{
		config: cfg,
	}

	slog

	if err := api.run(api.mount()); err != nil {
		log.Printf("Server failed to start ,err : %s", err)
		os.Exit(1)
	}
}
