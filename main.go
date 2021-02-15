package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	htmlParser "github.com/connectwithub/go-html-parser/html-parser"
)

func main() {
	local := flag.Bool("local", false, "Parse local html file (defaul false)")
	loc := flag.String("loc", "", "The location to parse links from")
	flag.Parse()

	startTime := time.Now()
	linkData := htmlParser.ParseHTMLLinks(*local, *loc)
	endTime := time.Since(startTime)

	fmt.Printf("Total Links: %v \n", len(linkData))
	fmt.Printf("Links: \n")
	fmt.Printf("--------------------------- \n")
	for _, link := range linkData {
		fmt.Printf("Link is: '%v' \n", link.Href)
		fmt.Printf("Text is: '%v' \n", link.Text)
		fmt.Printf("--------------------------- \n")
	}
	log.Printf("Parses took %s", endTime)
}
