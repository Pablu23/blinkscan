package main

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/pablu23/blinkscan/backend"
	"github.com/pablu23/blinkscan/backend/config"
	"github.com/pablu23/blinkscan/backend/database"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()
	config := config.FromEnv()

	//Todo: Do this in config not like this
	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	log.Debug().Interface("config", config).Msg("Loaded config")

	conn, err := pgx.Connect(ctx, config.PostgresConfig.ConnectionString())
	if err != nil {
		log.Fatal().Err(err).Msg("Could not connect to Postgres DB")
	}
	defer conn.Close(ctx)

	queries := database.New(conn)

	service := backend.NewService(queries)
	// pipeline := middleware.Pipeline(
	// 	middleware.RequestLogger,
	// )

	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	mux.HandleFunc("POST /api/account", service.PostAccount)
	mux.HandleFunc("POST /api/account/login", service.PostAccountLogin)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("Could not start http Server")
	}
}
