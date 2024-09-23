package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/pablu23/blinkscan/database"
	"github.com/rs/zerolog/log"
)

func Auth(db *database.Queries) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		log.Info().Msg("Enabling Auth")
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			sessionId := r.Header.Get("Authorization")
			id, err := uuid.Parse(sessionId)
			if err != nil {
				log.Warn().Err(err).Str("session-id", sessionId).Msg("Could not convert session-id to uuid")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			user, err := db.GetUserForSession(ctx, id)
			if err != nil {
				log.Warn().Err(err).Str("session-id", sessionId).Msg("No session for session-id")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			ctx = context.WithValue(ctx, "user", user)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
