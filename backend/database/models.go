// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	ID            pgtype.UUID
	Name          string
	Base64PwdHash string
	Base64PwdSalt string
}

type AccountSubscribedManga struct {
	ID        pgtype.UUID
	AccountID pgtype.UUID
	MangaID   pgtype.UUID
}

type AccountViewedChapter struct {
	ID        pgtype.UUID
	AccountID pgtype.UUID
	ChapterID pgtype.UUID
	ViewedAt  pgtype.Timestamp
}

type Chapter struct {
	ID      pgtype.UUID
	Title   string
	Number  int32
	MangaID pgtype.UUID
}

type ChapterImage struct {
	ChapterID pgtype.UUID
	ImageID   pgtype.UUID
	Alignment int32
}

type Image struct {
	ID   pgtype.UUID
	Path string
}

type Manga struct {
	ID            pgtype.UUID
	ProviderID    pgtype.UUID
	Title         string
	ThumbnailID   pgtype.UUID
	LatestChapter pgtype.Int4
	RequestedFrom pgtype.UUID
	LastUpdated   pgtype.Timestamp
	Created       pgtype.Timestamp
}

type Provider struct {
	ID   pgtype.UUID
	Url  string
	Name string
}

type Session struct {
	ID        pgtype.UUID
	AccountID pgtype.UUID
	Created   pgtype.Timestamp
}
