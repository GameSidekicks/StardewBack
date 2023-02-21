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
	Image    string `json:"image"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}

type Items struct {
	Item       string `json:"item"`
	Collection string `json:"collection"`
}

func main() {

	// allVillagers := make([]Villa, 0)

	// Init collectors
	cacheDir := filepath.Join("cache")

	// colVillager := colly.NewCollector(
	// 	colly.CacheDir(cacheDir),
	// 	colly.AllowedDomains("stardewvalleywiki.com"),
	// )

	// // Feedback: which URL are we scraping
	// colVillager.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Scraping:", r.URL)
	// })

	// // Feedback: response status
	// colVillager.OnResponse(func(r *colly.Response) {
	// 	fmt.Println("Status:", r.StatusCode)
	// })

	// // Error
	// colVillager.OnError(func(r *colly.Response, err error) {
	// 	fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	// })

	// colVillager.OnHTML(".mw-parser-output tr", func(t *colly.HTMLElement) {
	// 	villager := t.ChildText("td:nth-child(1)")
	// 	birth := t.ChildText("td:nth-child(2)")
	// 	photo := t.DOM.Find("td:nth-child(1) img").AttrOr("src", "none")

	// 	if villager != "" && villager != "Universals" {

	// 		dataTable := Villa{
	// 			Image:    "stardewvalleywiki.com" + photo,
	// 			Name:     villager,
	// 			Birthday: birth,
	// 		}
	// 		allVillagers = append(allVillagers, dataTable)
	// 	}

	// })

	// colVillager.Visit("https://stardewvalleywiki.com/List_of_All_Gifts")

	// writeJsonVilla(allVillagers)

	// The items collector

	allItems := make([]Items, 0)

	colItems := colly.NewCollector(
		colly.CacheDir(cacheDir),
		colly.AllowedDomains("stardewvalleywiki.com"),
	)

	// Feedback: which URL are we scraping
	colItems.OnRequest(func(r *colly.Request) {
		fmt.Println("Scraping:", r.URL)
	})

	// Feedback: response status
	colItems.OnResponse(func(r *colly.Response) {
		fmt.Println("Status:", r.StatusCode)
	})

	colItems.OnXML(itemsDivpath, func(t *colly.XMLElement) {
		item := t.ChildText(itemsNameXPath)
		// collection :=

		dataTableItems := Items{
			Item: item,
			// Collection: collection,
		}

		allItems = append(allItems, dataTableItems)

	})

	colItems.Visit("https://stardewvalleywiki.com/Collections")

	writeJsonItems(allItems)
}

func writeJsonVilla(data []Villa) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}

	_ = ioutil.WriteFile("villagers.json", file, 0644)
}

func writeJsonItems(data []Items) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}

	_ = ioutil.WriteFile("items_collection.json", file, 0644)
}
