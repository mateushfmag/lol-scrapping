package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type SeasonSummary struct {
	TeamName string
	Position int
	SeriesWon int
	SeriesLost int
	SeriesWinRate int
	TotalSeries int
	GamesWon int
	GamesLost int
	GamesWinRate int
	TotalGames int
	Streak int
}
func getResultsPerTeam(c *colly.Collector){
	c.OnHTML(".standings-outer-div table tbody tr.teamhighlight.teamhighlighter", func(htmlElement *colly.HTMLElement) {

		summary := SeasonSummary{}

		summary.TeamName = htmlElement.Attr("data-teamhighlight")
		summary.Position = Str2int(htmlElement.ChildText("td:nth-child(1)"))

		series := strings.Split(htmlElement.ChildText("td:nth-child(3)")," - ")
		summary.SeriesWon, summary.SeriesLost = Str2int(series[0]), Str2int(series[1])
		summary.SeriesWinRate = Str2int(strings.ReplaceAll(htmlElement.ChildText("td:nth-child(4)"),"%",""))
		summary.TotalSeries = summary.SeriesWon + summary.SeriesLost

		games := strings.Split(htmlElement.ChildText("td:nth-child(5)")," - ")
		summary.GamesWon, summary.GamesLost = Str2int(games[0]), Str2int(games[1])
		summary.GamesWinRate = Str2int(strings.ReplaceAll(htmlElement.ChildText("td:nth-child(6)"),"%",""))
		summary.TotalGames = summary.GamesWon + summary.GamesLost


		streak := htmlElement.ChildText("td:nth-child(8)")
		if strings.Contains(streak,"W") {
			summary.Streak = Str2int(strings.ReplaceAll(streak,"W",""))
		}else{
			summary.Streak = -1*Str2int(strings.ReplaceAll(streak,"L",""))
		}

		result, _ := json.Marshal(summary)
        fmt.Println(string(result))
    })
}

func GetLckResults(){
	c := colly.NewCollector()
	getResultsPerTeam(c)
	c.Visit("https://lol.fandom.com/wiki/LCK/2023_Season/Spring_Season")
}