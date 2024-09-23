// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	ID            uuid.UUID
	Name          string
	Base64PwdHash string
	Base64PwdSalt string
}

type AccountSubscribedManga struct {
	ID        uuid.UUID
	AccountID uuid.UUID
	MangaID   uuid.UUID
}

type AccountViewedChapter struct {
	ID        uuid.UUID
	AccountID uuid.UUID
	ChapterID uuid.UUID
	ViewedAt  pgtype.Timestamp
}

type Chapter struct {
	ID      uuid.UUID
	Title   string
	Number  int32
	Url     string
	MangaID uuid.UUID
}

type ChapterImage struct {
	ChapterID uuid.UUID
	ImageID   uuid.UUID
	Alignment int32
}

type Image struct {
	ID   uuid.UUID
	Path string
}

type Manga struct {
	ID            uuid.UUID
	ProviderID    uuid.UUID
	Title         string
	ThumbnailID   pgtype.UUID
	LatestChapter pgtype.Int4
	RequestedFrom pgtype.UUID
	LastUpdated   pgtype.Timestamp
	Created       pgtype.Timestamp
}

type Provider struct {
	ID   uuid.UUID
	Url  string
	Name string
}

type Session struct {
	ID        uuid.UUID
	AccountID uuid.UUID
	Created   pgtype.Timestamp
}
