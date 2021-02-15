package htmlParser

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"

	models "github.com/connectwithub/go-html-parser/Models"
)

//ParseHTMLLinks
func ParseHTMLLinks(local bool, loc string) []models.Link {
	var linkData = []models.Link{}
	switch local {
	case true:
		linkData = parseLocalHTML(loc)
	case false:
		linkData = parseOnlineHTML(loc)
	}
	return linkData
}

func parseLocalHTML(loc string) []models.Link {
	file, err := os.Open(loc)
	if err != nil {
		log.Fatalf("unable to read file %v, Error: %v", loc, err)
	}
	var fileReader io.Reader = file
	return parseLinks(&fileReader)
}

func parseOnlineHTML(loc string) []models.Link {
	resp, err := http.Get(loc)
	if err != nil {
		log.Fatalf("Unable to load page Error: %v", err)
	}
	var respData io.Reader = resp.Body
	return parseLinks(&respData)
}

func parseLinks(file *io.Reader) []models.Link {
	linkData := []models.Link{}
	doc, err := html.Parse(*file)
	if err != nil {
		log.Fatalf("Parse Error: %v", err)
	}
	var f func(*html.Node)
	if err != nil {
		log.Fatalf("Parse Error: %v", err)
	}
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					text := &bytes.Buffer{}
					collectText(n, text)
					linkData = append(linkData, models.Link{Href: a.Val, Text: strings.Trim(text.String(), " \n")})
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return linkData
}

func collectText(n *html.Node, buf *bytes.Buffer) {
	if n.Type == html.TextNode {
		buf.WriteString(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		collectText(c, buf)
	}
}
