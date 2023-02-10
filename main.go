package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
    c := colly.NewCollector(
        colly.AllowedDomains("en.wikipedia.org"),
    )
	c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
        links := e.ChildAttrs("a", "href")
		result := strings.Join(links,"\n")
        fmt.Println(result)
    })
    c.Visit("https://en.wikipedia.org/wiki/Web_scraping")
}