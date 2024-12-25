package provider

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
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

type ChapterResult struct {
	URL    Url
	Number int
	Title  string
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

	resp, err := http.Get(u.String())
	if err != nil {
		log.Error().Err(err).Str("url", u.String()).Str("provider", a.underlying.Name).Msg("Could not get Mangas")
		return nil, err
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var traverse func(*html.Node) []SearchResult
	traverse = func(n *html.Node) []SearchResult {
		if n.Type == html.ElementNode && n.Data == "div" {
			if checkDivForAttr(n, "grid grid-cols-2 sm:grid-cols-2 md:grid-cols-5 gap-3 p-4") {
				res := make([]SearchResult, 0)
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					href, err := extractHref(c)
					if err != nil {
						continue
					}

					imgNode := extractNested(c, "img")
					if imgNode == nil {
						continue
					}

					img, err := extractImg(imgNode)
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

func checkDivForAttr(n *html.Node, attr string) bool {
	for _, a := range n.Attr {
		if a.Key == "class" && strings.Contains(a.Val, attr) {
			return true
		}
	}
	return false
}

func extractNested(n *html.Node, data string) *html.Node {
	if n.Type == html.ElementNode && n.Data == data {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := extractNested(c, data)
		if node != nil {
			return node
		}
	}

	return nil
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

func (a *AsuraToons) GetChapters(manga database.Manga) []ChapterResult {
	mangaUrl, err := url.JoinPath(a.underlying.Url, "series", manga.InternalID+"-00000000")
	if err != nil {
		return nil
	}
	resp, err := http.Get(mangaUrl)
	if err != nil {
		log.Error().Err(err).Str("url", mangaUrl).Str("provider", a.underlying.Name).Str("manga_id", manga.ID.String()).Msg("Could not get Mangas")
		return nil
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic("")
	}

	res := make([]ChapterResult, 0)
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" {
			if checkDivForAttr(n, "pl-4 py-2 border rounded-md group w-full hover:bg-[#343434] cursor-pointer border-[#A2A2A2]/20 relative") {
				href, err := extractHref(n.FirstChild)
				if err != nil {
					return
				}
				curr := n.FirstChild

				var chapterName string = ""
				for c := curr.FirstChild; c != nil; c = c.NextSibling {
					if c.Data == "h3" && checkDivForAttr(c, "text-white") {
						spanNode := extractNested(c, "span")
						if spanNode != nil {
							if spanNode.FirstChild != nil {
								chapterName = spanNode.FirstChild.Data
							}
						}
					}
				}

				number := 0
				parts := strings.Split(string(href), "/")
				if len(parts) >= 0 {
					num := parts[len(parts)-1]
					number, _ = strconv.Atoi(num)
				}

				res = append(res, ChapterResult{
					URL:    href,
					Number: number,
					Title:  chapterName,
				})
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(doc)

	return res
}

func (a *AsuraToons) GetChapterImages(chapter database.Chapter) []database.ChapterImage {
	panic("Implement Me :P")
}
