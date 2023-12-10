package main

import (
	"fmt"
	"strings"

	"github.com/souravdey425/link"
)

var examplesHtml = `<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(examplesHtml)
	links, err := link.Parse(r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(links)
}
