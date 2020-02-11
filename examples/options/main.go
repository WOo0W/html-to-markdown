package main

import (
	"fmt"
	"log"

	md "github.com/WOo0W/html2md"
)

func main() {
	html := `<strong>Bold Text</strong>`
	// -> `__Bold Text__`
	// instead of `**Bold Text**`

	opt := &md.Options{
		StrongDelimiter: "__", // default: **
	}
	conv := md.NewConverter("", true, opt)

	markdown, err := conv.ConvertString(html)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(markdown)
}
