package provider

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"
	"github.com/pablu23/blinkscan/database"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/html"
)

type Url string

type SearchResult struct {
	URL       Url
	Thumbnail Url
}

type Provider interface {
	UUID() uuid.UUID
	URL() Url
	SearchMangas(name string) ([]SearchResult, error)
	GetChapters(manga database.Manga) ([]database.Chapter, error)
	GetChapterImages(chapter database.Chapter) ([]database.ChapterImage, error)
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

// Should probably not log inhere
func (a *AsuraToons) SearchMangas(name string) ([]SearchResult, error) {
	u, err := url.Parse(a.underlying.Url)
	if err != nil {
		log.Error().Err(err).Str("url", a.underlying.Url).Str("provider", a.underlying.Name).Msg("Could not parse Url")
		return nil, err
	}
	u = u.JoinPath("series")
	q := u.Query()
	q.Set("name", name)
	u.RawQuery = q.Encode()
	// fmt.Println(u.String())

	resp, err := http.Get(u.String())
	if err != nil {
		log.Error().Err(err).Str("url", u.String()).Str("provider", a.underlying.Name).Msg("Could not get Mangas")
		return nil, err
	}

	// html, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, err
	// }

	// _ = html
	// fmt.Println(string(html))
	// os.WriteFile("search.html", html, 0777)

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var traverse func(*html.Node) []SearchResult
	traverse = func(n *html.Node) []SearchResult {
		if n.Type == html.ElementNode && n.Data == "div" {
			if res := proccessDiv(n); res != nil {
				return res
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			res := traverse(c)
			if res != nil {
				return res
			}
		}
		return nil
	}
	res := traverse(doc)
	return res, nil
}

func proccessDiv(n *html.Node) []SearchResult {
	for _, a := range n.Attr {
		if a.Key == "class" && strings.Contains(a.Val, "grid grid-cols-2 sm:grid-cols-2 md:grid-cols-5 gap-3 p-4") {
			res := make([]SearchResult, 0)
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				href, err := extractHref(c)
				if err != nil {
					continue
				}

				img, err := extractNestedImg(c)
				if err != nil {
					continue
				}

				res = append(res, SearchResult{
					URL:       href,
					Thumbnail: img,
				})
			}
			return res
		}
	}
	return nil
}

func extractNestedImg(n *html.Node) (Url, error) {
	if n.Type == html.ElementNode && n.Data == "img" {
		img, err := extractImg(n)
		return img, err
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		img, err := extractNestedImg(c)
		if err == nil {
			return img, err
		}
	}
	return Url(""), errors.New("Could not find Image")
}

func extractImg(n *html.Node) (Url, error) {
	for _, a := range n.Attr {
		if a.Key == "src" {
			return Url(a.Val), nil
		}
	}

	return Url(""), errors.New("Could not find src")
}

func extractHref(n *html.Node) (Url, error) {
	for _, a := range n.Attr {
		if a.Key == "href" {
			return Url(a.Val), nil
		}
	}

	return Url(""), errors.New("Could not find href")
}

func (a *AsuraToons) GetChapters(manga database.Manga) []database.Chapter {
	panic("Implement Me :P")
}

func (a *AsuraToons) GetChapterImages(chapter database.Chapter) []database.ChapterImage {
	panic("Implement Me :P")
}
