# Autolycus - a universal web scrapper integration package

![logo](https://user-images.githubusercontent.com/66695611/236683354-6d57bbd0-46b6-4386-9ae3-3c03e55d7a0c.png)
In Greek Mythology Autolycus was a legendary thief, renowned by all...

Based on that, I present - Autolycus, a simple web scrapper package based on the [soup](https://github.com/anaskhan96/soup) module, to steal all the data you want... :)

* This package was specially made for [netBlast](https://github.com/Wikpi/Autolycus.git), to scrape colors for UI. (Examples for tthat can be found in `./example`)

# Features
**Clean layout**

**Fast**

**Easy to start**

# Installation / Usage
For an in-deoth look at the package, you can clone the repository:
```
git clone https://github.com/Wikpi/Autolycus.git
```

To use the Autolycus package, you need:
* Website URL to scrape
* Argument to scrape (tag, key, value)
* Write path to store all the data (optional)
* Action to take i.e. print, write (optional)

You can either use the all in one function:
```
data := autolycus.ScrapeData(scrapeURL, [tag, key, value], scrapePath, action)
```

Or have more control and do it all yourself. For that you first have to initiliaze the autolycus package:
```
doc := autolycus.Initiate(scrapeURL)
```

After which you scrape your data, to the desired slice variable:
```
Scrape(&data, doc, [tag, key, value])
```

When that is done, you can choose to either write or print out the data:
```
WriteData(scrapePath, data)
```
```
PrintData(data)
```


