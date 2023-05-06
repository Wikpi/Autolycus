package autolycus

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

// Variable, where the scrapped data is stored
var data []string

// Temporary scrape function
func ScrapeData(scrapeURL string, arg string, scrapePath string, action string) []string {
	c := Initiate(scrapeURL)
	GetRequest(c)

	Scrape(c, arg)

	switch action {
	case "give":
		return GiveData(c)
	case "print":
		PrintData(c)
	case "file":
		WriteData(c, scrapePath)
	}

	Visit(c, scrapeURL)

	return nil
}

// Initiates the scraper
func Initiate(scrapeURL string) *colly.Collector {
	domain := getDomain(scrapeURL)

	c := colly.NewCollector(colly.AllowedDomains(domain))

	return c
}

// Visits the provided URL
func Visit(c *colly.Collector, scrapeURL string) {
	c.Visit(scrapeURL)
}

// Finds the domain name from the given URL
func getDomain(scrapeURL string) string {
	domain := ""

	if strings.Contains(scrapeURL, "http") {
		domain = strings.Split(scrapeURL, "/")[2]
	} else {
		domain = strings.Split(scrapeURL, "/")[0]
	}

	return domain
}

// Scrapes the given args from the website
func Scrape(c *colly.Collector, arg string) []string {
	var lData []string

	c.OnHTML(arg, func(h *colly.HTMLElement) {
		lData = append(lData, h.Text)
		data = append(data, h.Text)
	})

	return lData
}

// Writes the scraped data to the specified filePath
func WriteData(c *colly.Collector, scrapePath string) {
	c.OnScraped(func(r *colly.Response) {
		file, err := os.OpenFile(scrapePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatal("Couldnt open scrapePath file: ", err)
		}
		defer file.Close()

		err = ioutil.WriteFile(scrapePath, []byte(""), 0644)

		for _, k := range data {
			_, err := file.WriteString(k + ", ")
			if err != nil {
				log.Fatal("Couldnt write to scrapePath file: ", err)
			}
		}
	})
}

// Writes the scraped data to the specified filePath
func PrintData(c *colly.Collector) {
	c.OnScraped(func(r *colly.Response) {
		fmt.Println(data)
	})
}

// Gives back data slice - not working!
func GiveData(c *colly.Collector) []string {
	var lData []string

	c.OnScraped(func(r *colly.Response) {
		lData = data
	})
	return lData
}

// Check for any errors while scraping the website
func CheckError(c *colly.Collector) {
	c.OnError(func(r *colly.Response, e error) {
		file, err := os.OpenFile("./logs/logs.txt", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			fmt.Print(err)
		}
		defer file.Close()

		// Writes error to logs file
		if _, err := file.WriteString(e.Error()); err != nil {
			fmt.Println(err)
		}

		// Exits program and gives message where error occured
		log.Fatal("Error scraping the website.")
	})
}

// Displays status
func GetRequest(c *colly.Collector) {
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL)
	})
}
