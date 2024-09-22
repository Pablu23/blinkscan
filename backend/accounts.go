package backend

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/pablu23/blinkscan/backend/database"
	"github.com/pablu23/blinkscan/backend/transport"
	"github.com/rs/zerolog/log"
)

func createPwdHash(pwd []byte, salt []byte) [32]byte {
	return sha256.Sum256(append(pwd, salt...))
}

func (s *Service) PostAccount(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var account transport.Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		log.Error().Err(err).Msg("Could not decode account json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	salt := make([]byte, 24)
	rand.Read(salt)
	hash := createPwdHash([]byte(account.Password), salt)
	b64hash := base64.StdEncoding.EncodeToString(hash[:])
	b64salt := base64.StdEncoding.EncodeToString(salt)

	_, err = s.db.GetAccountByName(ctx, account.Username)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		log.Error().Err(err).Str("name", account.Username).Msg("Could not get account by name")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if err == nil {
		log.Warn().Str("name", account.Username).Msg("User by name already exists")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Info().Str("name", account.Username).Msg("Created account")
	s.db.CreateAccount(ctx, database.CreateAccountParams{
		Name:          account.Username,
		Base64PwdHash: b64hash,
		Base64PwdSalt: b64salt,
	})
}

func (s *Service) PostAccountLogin(w http.ResponseWriter, r *http.Request) {
	// TODO: Jwt token
	ctx := context.Background()
	var loginCredentials transport.Account
	err := json.NewDecoder(r.Body).Decode(&loginCredentials)
	if err != nil {
		log.Error().Err(err).Msg("Could not decode account json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	acc, err := s.db.GetAccountByName(ctx, loginCredentials.Username)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		log.Error().Err(err).Str("name", loginCredentials.Username).Msg("Could not get account by name")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if err != nil {
		log.Warn().Str("name", loginCredentials.Username).Msg("User by name doesnt exists")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	salt, err := base64.StdEncoding.DecodeString(acc.Base64PwdSalt)
	if err != nil {
		log.Error().Err(err).Msg("Could not decode b64 salt")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	expectedHash, err := base64.StdEncoding.DecodeString(acc.Base64PwdHash)
	if err != nil {
		log.Error().Err(err).Msg("Could not decode b64 hash")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	hash := createPwdHash([]byte(loginCredentials.Password), salt)

	if bytes.Equal(hash[:], expectedHash) {
		log.Trace().Str("username", loginCredentials.Username).Msg("Login succeded")
		w.WriteHeader(http.StatusOK)
		session, err := s.db.CreateSession(ctx, acc.ID)
		if err != nil {
			log.Error().Err(err).Str("account", acc.ID.String()).Msg("Could not create session")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//TODO: Dont use uuid, use cryptographically secure "fingerprint"
		w.Write([]byte(session.ID.String()))
		//TODO: return jwt token
	} else {
		log.Debug().Str("username", loginCredentials.Username).Msg("Login failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
