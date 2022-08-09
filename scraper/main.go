package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Weather struct {
   Week     		string
   Temperature    	string
   Url    			string
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("weather.gc.ca", "www.weather.gc.ca"),
	)
	
	weathers := make([]Weather, 0)


	c.OnHTML("div.div-table", func(e *colly.HTMLElement) {
      e.ForEach("div.div-column", func(i int, h *colly.HTMLElement) {
           item := Weather{}
           item.Week = h.ChildText("[class='div-row div-row1 div-row-head']")
           item.Temperature = h.ChildText("span.high")[:2]
           item.Url = e.ChildAttrs("a", "href")[1]
           weathers = append(weathers, item)
       })
	})
	
   c.OnScraped(func(r *colly.Response) {
       fmt.Println("Finished", r.Request.URL)
       js, err := json.MarshalIndent(weathers, "", "    ")
       if err != nil {
           log.Fatal(err)
       }
       fmt.Println("Writing data to file")
       if err := os.WriteFile("forecasts.json", js, 0664); err == nil {
           fmt.Println("Data written to file successfully")
       }

   })

	c.Visit("https://weather.gc.ca/city/pages/qc-147_metric_e.html")
}