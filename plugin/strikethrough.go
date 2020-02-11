package plugin

import (
	"strings"

	"github.com/WOo0W/html2md"

	"github.com/PuerkitoBio/goquery"
)

// Strikethrough converts `<strike>`, `<s>`, and `<del>` elements
func Strikethrough(character string) html2md.Plugin {
	return func(c *html2md.Converter) []html2md.Rule {
		if character == "" {
			character = "~"
		}

		return []html2md.Rule{
			html2md.Rule{
				Filter: []string{"del", "s", "strike"},
				Replacement: func(content string, selec *goquery.Selection, opt *html2md.Options) *string {
					// trim spaces so that the following does NOT happen: `~ and cake~`
					content = strings.TrimSpace(content)
					return html2md.String(character + content + character)
				},
			},
		}
	}
}
