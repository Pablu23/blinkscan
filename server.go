package blinkscan

import "github.com/pablu23/blinkscan/database"

type Service struct {
	db *database.Queries
}

func NewService(db *database.Queries) *Service {
	return &Service{
		db: db,
	}
}
