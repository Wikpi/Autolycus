package main

import (
	autolycus "github.com/Wikpi/Autolycus/pkg"
)

func main() {
	// Temporary
	autolycus.ScrapeData("https://htmlcolorcodes.com/colors/", "td.color-table__cell--hex", "./build/data.txt", "file")

	// // Creates a new collector from the colly library
	// c := autolycus.Initiate("https://htmlcolorcodes.com/colors/")
	// autolycus.GetRequest(c)
	// autolycus.CheckError(c)

	// //data := autolycus.Scrape(c, "td.color-table__cell--hex")

	// fmt.Println(c)
	// autolycus.PrintData(c)

	// autolycus.WriteData(c, "./build/data.txt")

	// autolycus.Visit(c, "https://htmlcolorcodes.com/colors/")
}
