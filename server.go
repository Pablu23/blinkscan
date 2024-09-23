package backend

import "github.com/pablu23/blinkscan/backend/database"

type Service struct {
	db *database.Queries
}

func NewService(db *database.Queries) *Service {
	return &Service{
		db: db,
	}
}
