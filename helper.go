package blinkscan

import (
	"context"
	"errors"

	"github.com/pablu23/blinkscan/database"
	"github.com/rs/zerolog/log"
)

// Gets Account from Context value, is only save on private Endpoints that were already authenticated
// If Context does not have a value of type database.Account for key "user" this function will return a default database.Account and an error
func GetAccount(ctx context.Context) (database.Account, error) {
	val := ctx.Value("user")
	switch val.(type) {
	case database.Account:
		return val.(database.Account), nil
	default:
		return database.Account{}, errors.New("Could not find Account in context")
	}
}

// Gets Account from Context value, is only save on private Endpoints that were already authenticated
// If Context does not have a value of type database.Account for key "user" this function will panic
func MustGetAccount(ctx context.Context) database.Account {
	account, err := GetAccount(ctx)
	if err != nil {
		log.Panic().Err(err).Send()
	}
	return account
}
