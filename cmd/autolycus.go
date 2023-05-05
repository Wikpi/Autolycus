package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

// go run ./pkg/autolycus.go -url https://htmlcolorcodes.com/colors/ -tag td.color-table__cell--hex -path ./build/data.txt

func main() {
	// Variable, where the scrapped data is stored
	var data []string

	scrapeURL, parameter, scrapePath := setFlags()

	domain := getDomain(scrapeURL)

	c := colly.NewCollector(colly.AllowedDomains(domain))

	c.OnHTML(parameter, func(h *colly.HTMLElement) {
		scrapeData(parameter, &data, h.Text)
	})

	c.OnError(func(r *colly.Response, e error) {
		checkError(e)
	})

	c.OnScraped(func(r *colly.Response) {
		writeData(scrapePath, data)
	})

	c.Visit(scrapeURL)
}

// Sets flgas, which have to be specified in order for the scrape to be successful
func setFlags() (string, string, string) {
	scrapeURL := flag.String("url", "", "Specify the URL to scrape")
	parameter := flag.String("tag", "", "Provide paramters to scrape, seperate tags seperated by a whitespace")
	scrapePath := flag.String("path", "", "Provide a suitable path to write the scraped data to")

	flag.Parse()

	if *scrapeURL == "" {
		log.Fatal("Didnt provide a URL to scrape")
	}
	if *parameter == "" {
		log.Fatal("Didnt provide a paramtere to scrpae")
	}
	if *scrapePath == "" {
		log.Fatal("Didnt provide a path to store scraped data")
	}

	return *scrapeURL, *parameter, *scrapePath
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

// Scrapes the given parameters from the website
func scrapeData(parameter string, data *[]string, value string) {
	*data = append(*data, value)
}

// Check for any errors while scraping the website
func checkError(e error) {
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
}

// Writes the scraped data to the specified filePath
func writeData(scrapePath string, data []string) {
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
}

// Displays status
func getRequest(c *colly.Collector) {
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL)
	})
}
