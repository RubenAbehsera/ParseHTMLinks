package main

import (
	"fmt"
	"link"
	"strings"
)

var exampleHTML = `
<html>
<body>
<h1>Hello!</h1>
<a href="/other-page">A link to another page</a>
<a href="/other-page2">A link to another page <span>with span</span></a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHTML)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n\n", links)
}
