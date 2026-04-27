package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/pauljomy/social/internal/env"
)

func main() {

	godotenv.Load()

	addr := env.GetString("ADDR", ":8080")
	appEnv := env.GetString("ENV", "development")

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cfg := config{
		addr: addr,
		env:  appEnv,
	}

	app := &application{
		logger: logger,
		config: cfg,
	}

	mount := app.mount()

	logger.Info("API server started at", "addr", addr)

	log.Fatal(app.run(mount))

}
