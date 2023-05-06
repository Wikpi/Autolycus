// Example2 one single function to scrape the data i.e. all color hexcodes from the website

package main

import (
	"fmt"

	autolycus "github.com/Wikpi/Autolycus/pkg"
)

func main() {
	// website to scrape
	url := "https://htmlcolorcodes.com/colors/"
	// Arguments to scrape (tag, key, value !)
	arg := []string{"td", "class", "color-table__cell--hex"}
	// Write path of the txt file
	path := "./build/data.txt"

	// Using one function
	data := autolycus.ScrapeData(url, arg, path, "")
	fmt.Println(data)
}
