package blinkscan

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

func (s *Service) RegisterPublicRoutes(mux *http.ServeMux) {
  log.Debug().Msg("Registered public routes")
	mux.HandleFunc("POST /account", s.PostAccount)
	mux.HandleFunc("POST /account/login", s.PostAccountLogin)
}

func (s *Service) RegisterPrivateRoutes(mux *http.ServeMux) {
  log.Debug().Msg("Registered private routes")
	mux.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := MustGetAccount(ctx)
		w.Write([]byte(fmt.Sprintf("Hello %s", user.Name)))
	})
}
