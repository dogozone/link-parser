package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	s := `<html>
	<body>
		<h1>Hello!</h1>
		<a href="/dog">
						<span>Something in a span</span>
						Text not in a span
						<!-- nNOT INCL -->
						<b>Bold text!</b>
					</a>
		<a href="/other-page">A link to another page</a>
	</body>
	</html>
	`
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}

	parseNode(doc)
	dfs(doc, " ")
	links := parseNode(doc)
	fmt.Printf("final %v\n", links)
}

type Link struct {
	Href, Text string
}

func parseNode(doc *html.Node) []Link {
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
func dfs(n *html.Node, padding string) {
	fmt.Println(padding, n.Data)

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
