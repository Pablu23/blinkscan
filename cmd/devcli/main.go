package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pablu23/blinkscan/database"
	"github.com/pablu23/blinkscan/provider"
)

func main() {
	// testSearchMangas()
	// testGetChapters()
	testImages()
}

func testImages() {
	chapter := database.Chapter{
		Url: "somebody-stop-the-pope-9df9d5f9/chapter/22",
	}

	provider.AsuraToon.GetChapterImages(chapter)
}

func testSearchMangas() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Search: ")
		text, _ := reader.ReadString('\n')
		mangas, err := provider.AsuraToon.SearchMangas(text)
		if err != nil {
			panic(err)
		}

		for _, res := range mangas {
			fmt.Printf("Link: %s\nImg: %s\n--------\n", res.URL, res.Thumbnail)
		}
	}
}

func testGetChapters() {
	manga := database.Manga{
		InternalID: "somebody-stop-the-pope",
	}

	chapters := provider.AsuraToon.GetChapters(manga)
	for _, chapter := range chapters {
		fmt.Println(chapter)
	}
}
