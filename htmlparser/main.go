package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type htmlCustom struct {
	href     string
	textspan string
}

func main() {
	htmlContsruct := read()

	fmt.Println(htmlContsruct)
}

func read() []htmlCustom {
	s := `<html>
	<body>
	  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
	</body>
	</html>
  `

	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}

	var f func(*html.Node)

	var htmlStruct []htmlCustom

	f = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, a := range node.Attr {
				if a.Key == "href" {
					htmlStruct = append(htmlStruct, htmlCustom{
						href:     a.Val,
						textspan: getAlltheText(node),
					})

					//Call a recursive function to get all the nodes text under the a tag

				}
			}
			//Now the time is to get the value of the span as we expect that to be there
			//fmt.Println(node.FirstChild.Data)
			fmt.Println()
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	return htmlStruct
}

func getAlltheText(n *html.Node) string {

	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {

		ret += getAlltheText(c)
		//fmt.Println(ret)
	}

	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	return strings.Join(strings.Fields(ret), " ")

}
