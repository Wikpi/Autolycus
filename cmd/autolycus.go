package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/anaskhan96/soup"
)

// Example
// go run autolycus.go -url https://htmlcolorcodes.com/colors/ -tag td -key class -value color-table__cell--hex -path ./build/data.txt

func main() {
	// Variable, where the scrapped data is stored
	var data []string

	scrapeURL, arg, scrapePath := setFlags()

	doc := Initiate(scrapeURL)

	scrape(&data, doc, arg)

	action := "print"

	switch action {
	case "write":
		writeData(scrapePath, data)
	case "print":
		printData(data)
	}
}

// Sets flags, which have to be specified in order for the scrape to be successful
func setFlags() (string, []string, string) {
	scrapeURL := flag.String("url", "", "Specify the URL to scrape")
	tag := flag.String("tag", "", "Tag to scrape")
	key := flag.String("key", "", "Key to scrape")
	value := flag.String("value", "", "Value to scrape")
	scrapePath := flag.String("path", "", "Provide a suitable path to write the scraped data to")

	flag.Parse()

	if *scrapeURL == "" {
		log.Fatal("Didnt provide a URL to scrape")
	}
	if *tag == "" {
		log.Fatal("Didnt provide a tag to scrape")
	}
	if *key == "" {
		log.Fatal("Didnt provide a key to scrape")
	}
	if *value == "" {
		log.Fatal("Didnt provide the value of the key to scrape")
	}
	if *scrapePath == "" {
		log.Fatal("Didnt provide a path to store scraped data")
	}

	return *scrapeURL, []string{*tag, *key, *value}, *scrapePath
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
func scrape(lData *[]string, doc soup.Root, arg []string) {
	colors := doc.FindAll(arg[0], arg[1], arg[2])

	for _, color := range colors {
		*lData = append(*lData, color.Text())
	}
}

// Writes the scraped data to the specified filePath
func writeData(scrapePath string, lData []string) {
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
func printData(lData []string) {
	fmt.Println(lData)
}
