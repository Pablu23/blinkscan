package provider

import (
	"net/http"
	"net/url"

	"github.com/google/uuid"
	"github.com/pablu23/blinkscan/backend/database"
	"github.com/rs/zerolog/log"
)

type Url string

type Provider interface {
	UUID() uuid.UUID
	URL() Url
	SearchMangas(name string) []database.Manga
	GetChapters(manga database.Manga) []database.Chapter
	GetChapterImages(chapter database.Chapter) []database.ChapterImage
}

type AsuraToons struct {
	underlying database.Provider
}

var AsuraToon = AsuraToons{
	underlying: database.Provider{
		ID:   uuid.MustParse("4254feb5-a362-46fb-97c4-47705930d858"),
		Url:  "https://asuracomic.net/",
		Name: "Asuracomic",
	},
}

func (a *AsuraToons) UUID() uuid.UUID {
	return a.underlying.ID
}

func (a *AsuraToons) URL() Url {
	return Url(a.underlying.Url)
}

func (a *AsuraToons) SearchMangas(name string) []database.Manga {
	u, err := url.Parse(a.underlying.Url)
	if err != nil {
		log.Error().Err(err).Str("url", a.underlying.Url).Str("provider", a.underlying.Name).Msg("Could not parse Url")
		return nil
	}
	u.JoinPath("series")
	u.Query().Add("name", name)

	resp, err := http.Get(u.String())
	if err != nil {
		log.Error().Err(err).Str("url", u.String()).Str("provider", a.underlying.Name).Msg("Could not get Mangas")
		return nil
	}

	panic("Implement Me :P")
}

func (a *AsuraToons) GetChapters(manga database.Manga) []database.Chapter {
	panic("Implement Me :P")
}

func (a *AsuraToons) GetChapterImages(chapter database.Chapter) []database.ChapterImage {
	panic("Implement Me :P")
}
