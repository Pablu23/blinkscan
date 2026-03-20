package provider

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/pablu23/blinkscan/database"
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
