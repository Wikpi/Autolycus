package autolycus

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/anaskhan96/soup"
)

// Variable, where the scrapped data is stored
var Data []string

// ONE FUNCTION TO SCRAPE IT ALL (Lotr)
func ScrapeData(scrapeURL string, arg []string, scrapePath string, actions ...string) *[]string {
	doc := Initiate(scrapeURL)

	Scrape(&Data, doc, arg)

	for _, action := range actions {
		switch action {
		case "write":
			WriteData(scrapePath, Data)
		case "print":
			PrintData(Data)
		}
	}

	return &Data
}

// Initiates the scraper (gets an html string and parses it)
func Initiate(scrapeURL string) soup.Root {
	resp, err := soup.Get(scrapeURL)
	if err != nil {
		log.Fatal("Couldnt get html string")
	}

	doc := soup.HTMLParse(resp)

	return doc
}

// Scrapes the provided argument (tag, key, value)
func Scrape(lData *[]string, doc soup.Root, arg []string) {
	colors := doc.FindAll(arg[0], arg[1], arg[2])

	for _, color := range colors {
		*lData = append(*lData, color.Text())
	}
}

// Writes the scraped data to the specified filePath
func WriteData(scrapePath string, lData []string) {
	file, err := os.OpenFile(scrapePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal("Couldnt open scrapePath file: ", err)
	}
	defer file.Close()

	err = ioutil.WriteFile(scrapePath, []byte(""), 0644)

	for _, k := range lData {
		_, err := file.WriteString(k + ", ")
		if err != nil {
			log.Fatal("Couldnt write to scrapePath file: ", err)
		}
	}
}

// Prints out the data
func PrintData(lData []string) {
	fmt.Println(lData)
}
