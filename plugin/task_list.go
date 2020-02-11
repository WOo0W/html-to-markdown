package plugin

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/WOo0W/html2md"
)

// TaskListItems converts checkboxes into task list items.
func TaskListItems() html2md.Plugin {
	return func(c *html2md.Converter) []html2md.Rule {
		return []html2md.Rule{
			html2md.Rule{
				Filter: []string{"input"},
				Replacement: func(content string, selec *goquery.Selection, opt *html2md.Options) *string {
					if !selec.Parent().Is("li") {
						return nil
					}
					if selec.AttrOr("type", "") != "checkbox" {
						return nil
					}

					_, ok := selec.Attr("checked")
					if ok {
						return html2md.String("[x] ")
					}
					return html2md.String("[ ] ")
				},
			},
		}
	}
}
