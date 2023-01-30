package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/gocolly/colly"
)

type Villa struct {
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}

func main() {

	allVillagers := make([]Villa, 0)

	// Init collectors
	cacheDir := filepath.Join("cache")

	collVillager := colly.NewCollector(
		colly.CacheDir(cacheDir),
		colly.AllowedDomains("stardewvalleywiki.com"),
	)

	// Feedback: which URL are we scraping
	collVillager.OnRequest(func(r *colly.Request) {
		fmt.Println("Scraping:", r.URL)
	})

	// Feedback: response status
	collVillager.OnResponse(func(r *colly.Response) {
		fmt.Println("Status:", r.StatusCode)
	})

	// Error
	collVillager.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	collVillager.OnHTML(".mw-parser-output tr", func(t *colly.HTMLElement) {
		villager := t.ChildText("td:nth-child(1)")
		birth := t.ChildText("td:nth-child(2)")

		if villager != "" && villager != "Universals" {

			dataTable := Villa{
				Name:     villager,
				Birthday: birth,
			}
			allVillagers = append(allVillagers, dataTable)
		}

	})

	collVillager.Visit("https://stardewvalleywiki.com/List_of_All_Gifts")

	// enc := json.NewEncoder(os.Stdout)
	// enc.SetIndent("", " ")
	// enc.Encode(allVillagers)

	writeJson(allVillagers)
}

func writeJson(data []Villa) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}

	_ = ioutil.WriteFile("villagers.json", file, 0644)
}
