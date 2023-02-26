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

	links := link.Parse(doc)
	fmt.Printf("final %v\n", links)
}
