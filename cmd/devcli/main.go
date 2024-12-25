package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pablu23/blinkscan/provider"
)

func main() {
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
