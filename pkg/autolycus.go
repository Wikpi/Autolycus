package autolycus

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func Initiate(scrapeURL string, parameter string, scrapePath string, showData ...bool) {
	// Variable, where the scrapped data is stored
	var data []string

	domain := getDomain(scrapeURL)

	c := colly.NewCollector(colly.AllowedDomains(domain))

	if showData[0] {
		fmt.Println(scrapeURL, " ", parameter, " ", scrapePath)
		getRequest(c)
	}

	c.OnHTML(parameter, func(h *colly.HTMLElement) {
		scrapeData(parameter, &data, h.Text)
	})

	c.OnError(func(r *colly.Response, e error) {
		checkError(e)
	})

	c.OnScraped(func(r *colly.Response) {
		if showData[0] {
			fmt.Println(data)
		}

		writeData(scrapePath, data)
	})

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
