package main

import (
	"fmt"
	"log"

	"github.com/WOo0W/html2md"
	"github.com/WOo0W/html2md/plugin"
)

func main() {
	html := `
	<ul>
		<li><input type=checkbox checked>Checked!</li>
		<li><input type=checkbox>Check Me!</li>
	</ul>
	`
	/*
		- [x] Checked!
		- [ ] Check Me!
	*/

	conv := html2md.NewConverter("", true, nil)

	// Use the `GitHubFlavored` plugin from the `plugin` package.
	conv.Use(plugin.GitHubFlavored())

	markdown, err := conv.ConvertString(html)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(markdown)
}
