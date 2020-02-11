// Package plugin contains all the rules that are not
// part of Commonmark like GitHub Flavored Markdown.
package plugin

import "github.com/WOo0W/html2md"

// GitHubFlavored is GitHub's Flavored Markdown
func GitHubFlavored() html2md.Plugin {
	return func(c *html2md.Converter) (rules []html2md.Rule) {
		rules = append(rules, Strikethrough("")(c)...)
		rules = append(rules, EXPERIMENTAL_Table...)
		rules = append(rules, TaskListItems()(c)...)
		return
	}
}
