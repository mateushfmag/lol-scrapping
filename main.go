package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
    c := colly.NewCollector()
    fName := "data.csv"
    file, err := os.Create(fName)
    if err != nil {
        log.Fatalf("Could not create file, err: %q", err)
        return
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

	LinksCollector(c)
	TableCollector(c,writer)
}