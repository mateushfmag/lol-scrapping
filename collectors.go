package main

import (
	"encoding/csv"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func LinksCollector(c *colly.Collector) {
	c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
        links := e.ChildAttrs("a", "href")
		result := strings.Join(links,"\n")
        fmt.Println(result)
    })
	c.Visit("https://en.wikipedia.org/wiki/Web_scraping")
}

func TableCollector(c *colly.Collector, writer *csv.Writer) {
	c.OnHTML("table#customers", func(e *colly.HTMLElement) {
        e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
            writer.Write([]string{
                el.ChildText("td:nth-child(1)"),
                el.ChildText("td:nth-child(2)"),
                el.ChildText("td:nth-child(3)"),
            })
        })
        fmt.Println("Scrapping Complete")
    })
    c.Visit("https://www.w3schools.com/html/html_tables.asp")

}