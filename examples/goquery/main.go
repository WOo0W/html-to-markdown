package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/WOo0W/html2md"
)

func main() {
	url := "https://blog.golang.org/godoc-documenting-go-code"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	content := doc.Find("#content")

	conv := html2md.NewConverter(html2md.DomainFromURL(url), true, nil)
	markdown := conv.Convert(content)

	fmt.Println(markdown)
}
