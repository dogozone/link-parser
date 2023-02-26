package main

import (
	"fmt"
	"log"
	"strings"

	linkParser "github.com/dogozone/link-parser"
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
	reader := strings.NewReader(s)

	links, err := linkParser.Parse(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("final %v\n", links)
}
