package main

import (
	"fmt"
	"log"

	"github.com/WOo0W/html2md"
)

func main() {
	html := `<strong>Bold Text</strong>`
	// -> `__Bold Text__`
	// instead of `**Bold Text**`

	opt := &html2md.Options{
		StrongDelimiter: "__", // default: **
	}
	conv := html2md.NewConverter("", true, opt)

	markdown, err := conv.ConvertString(html)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(markdown)
}
