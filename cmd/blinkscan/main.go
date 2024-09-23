package main

import (
	"context"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/pablu23/blinkscan"
	"github.com/pablu23/blinkscan/config"
	"github.com/pablu23/blinkscan/database"
	"github.com/pablu23/blinkscan/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()
	config := config.FromEnv()

	//Todo: Do this in config not like this
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Debug().Interface("config", config).Msg("Loaded config")

	conn, err := pgx.Connect(ctx, config.PostgresConfig.ConnectionString())
	if err != nil {
		log.Fatal().Err(err).Msg("Could not connect to Postgres DB")
	}
	defer conn.Close(ctx)

	queries := database.New(conn)
	service := blinkscan.NewService(queries)

	publicMux := http.NewServeMux()
	service.RegisterPublicRoutes(publicMux)

	privateMux := http.NewServeMux()
	service.RegisterPrivateRoutes(privateMux)

	auth := middleware.Auth(queries)
	frontend := http.FileServer(http.Dir("./frontend/browser"))

	rootMux := http.NewServeMux()
	rootMux.Handle("/", frontend)
	rootMux.Handle("/api/", http.StripPrefix("/api", auth(privateMux)))
	rootMux.Handle("/api/public/", http.StripPrefix("/api/public", publicMux))
	rootMux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: rootMux,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("Could not start http Server")
	}
}
