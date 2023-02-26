package linkParser

import (
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href, Text string
}

func Parse(doc *html.Node) []Link {
	var links []Link
	var f func(*html.Node)
	f = func(n *html.Node) {

		if n.Type == html.ElementNode && n.Data == "a" {
			link := buildLink(n)
			links = append(links, link)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	return links
}

func buildLink(node *html.Node) Link {
	var link Link
	for _, attr := range node.Attr {
		if attr.Key == "href" {
			link.Href = attr.Val
			break
		}
	}
	link.Text = text(node)
	return link
}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}

	var builtText string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		builtText += text(c) + " "
	}

	// Using (strings) Fields to remove whitespace characters
	return strings.Join(strings.Fields(builtText), " ")
}

func extractText(ahref *html.Node) []string {
	var text []string
	var f func(*html.Node)
	f = func(n *html.Node) {

		if n.Type == html.TextNode {
			trimmed := strings.Join(strings.Fields(n.Data), " ")
			text = append(text, trimmed)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(ahref)

	return text
}

// Basic Depth-first search
// func dfs(n *html.Node, padding string) {
// 	fmt.Println(padding, n.Data)

// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		dfs(c, padding+"  ")
// 	}
// }
