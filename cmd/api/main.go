package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/pauljomy/social/internal/env"
	"github.com/pauljomy/social/internal/store/postgres"
)

func main() {

	godotenv.Load()

	addr := env.GetString("ADDR", ":8080")
	appEnv := env.GetString("ENV", "development")

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cfg := config{
		addr: addr,
		env:  appEnv,
		db: dbConfig{
			dsn:          env.GetString("DB_DSN", "postgres://user:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 25),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 25),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	posts := postgres.NewPostStore(nil)

	app := &application{
		logger: logger,
		config: cfg,
		posts:  posts,
	}

	mount := app.mount()

	logger.Info("API server started at", "addr", addr)

	log.Fatal(app.run(mount))

}
