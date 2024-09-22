package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/pablu23/blinkscan/backend"
	"github.com/pablu23/blinkscan/backend/config"
	"github.com/pablu23/blinkscan/backend/database"
	"github.com/pablu23/blinkscan/backend/middleware"
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
	service := backend.NewService(queries)

	publicMux := http.NewServeMux()
	publicMux.HandleFunc("POST /account", service.PostAccount)
	publicMux.HandleFunc("POST /account/login", service.PostAccountLogin)

	privateMux := http.NewServeMux()
	privateMux.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(database.Account)
		w.Write([]byte(fmt.Sprintf("Hello %s", user.Name)))
	})

	auth := middleware.Auth(queries)
	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", auth(privateMux)))
	mux.Handle("/", publicMux)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("Could not start http Server")
	}
}
