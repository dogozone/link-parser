package linkParser

import (
	"log"
	"strings"
	"testing"
)

var S1 string = `<html>
<body>
	<a href="/page">A link to a page</a>
</body>
</html>
`

func TestParse(t *testing.T) {
	reader := strings.NewReader(S1)

	links, err := Parse(reader)
	if err != nil {
		log.Fatal(err)
	}
	expected := Link{
		Href: "/page",
		Text: "A link to a page",
	}
	if expected.Text != links[0].Text {
		t.Errorf("parse link Text: \nexpected: %s,\nactual: %s\n", expected.Text, links[0].Text)
	}
}
