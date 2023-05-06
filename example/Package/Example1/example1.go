// Example1 showing different functions to scrape the data i.e. all color hexcodes from the website

package main

import (
	autolycus "github.com/Wikpi/Autolycus/pkg"
)

func main() {
	colors := []string{}
	// website to scrape
	url := "https://htmlcolorcodes.com/colors/"
	// Arguments to scrape (tag, key, value !)
	arg := []string{"td", "class", "color-table__cell--hex"}
	// Write path of the txt file
	path := "./build/data.txt"

	// Iniates the scrapper i.e. get the html string and parses it
	doc := autolycus.Initiate(url)
	// Scrapes the data
	autolycus.Scrape(&colors, doc, arg)

	// Determines what to do
	action := "write"

	switch action {
	case "write":
		autolycus.WriteData(path, colors)
	case "print":
		autolycus.PrintData(colors)
	}
}
