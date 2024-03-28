 package main

 import (
   "fmt"
   "log"
   "strings"

  "github.com/gocolly/colly/v2"
 )

 func main() {
   //Create a new Colly collector
   c := colly.NewCollector()

   //Define the URL you want to scrape
   url := "http://quotes.toscrape.com/page/1/"

   //Set up callbacks to handle scrapping events
   c.OnHTML(".quote", func(e *colly.HTMLElement) {
     //Extract data from html elements
     quote := e.ChildText("span.text")
     author := e.ChildText("small.author")
     tags := e.ChildText("div.tags")

     //Clean up the extracted data
     quote = strings.TrimSpace(quote)
     author = strings.TrimSpace(author)
     tags = strings.TrimSpace(tags)

     //Print the scraped data
     fmt.Printf("Quote: %s\nAuthor: %s\nTags: %s\n\n", quote, author, tags)
   })

   //Visit the URL and start scrapping
   err := c.Visit(url)
   if err != nil {
     log.Fatal(err)
   }
 }
