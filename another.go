 package main

 import (
   "fmt"
   "github.com/gocolly/colly"
   "encoding/csv"
   "os"
   "log"
 )

 func main(){
   fmt.Println("Hello, world!")

   c := colly.NewCollector()
   c.Visit("https://en.wikipedia.org/wiki/Main_Page")

   c.OnRequest(func(e *colly.Request){
     fmt.Println("Visiting: ", err)
   })

   c.OnError(func(_ *colly.Response){
     log.Println("Something went wrong: ", err)
   })

   c.OnResponse(func(r *colly.Response){
     fmt.Println("Page visited: ", r.Request.URL)
   })

   c.OnHTML("a", func(e *colly.HTMLElement){
     fmt.Println("%v", e.Attr("href"))
   })

   c.OnScraped(func(e *colly.Response){
     fmt.Println(r.Request.URL, " scraped!")
   })

   c.Visit("https://scrapeme.live/shop/")

   type PokemonProduct struct {
     url, image, name, price, string
   }

   var pokemonProducts []PokemonProduct

   c.OnHTML("li.product", func(e *colly.HTMLElement){
     //Initializing a new PokemonProduct instance
     pokemonProduct := PokemonProduct{}

     //Scraping the data of interests
     pokemonProduct.url = e.ChildAttr("a", "href")
     pokemonProduct.image = e.ChildAttr("img", "src")
     pokemonProduct.name = e.ChildAttr("h2")
     pokemonProduct.price = e.ChildAttr(".price")

     //Adding the product instance with scraped data to the list of products 
     pokemonProducts = append(pokemonProducts, pokemonProduct)
   })


  // Opening CSV file
   file, err := os.Create("producst.csv")
   ir err != nil {
     log.Fatalln("Failed to create output CSV file.", err)
   }
   defer file.Close()

   //Initializing a file writer
   writer := csv.NewWriter(file)

   //Defining the CSV headers
   headers := []string{
     "url"
     "image"
     "name"
     "price"
   }

   //Writing the column headers
   writer.Write(headers)

   //adding each Polemon product to the CSV output file
   for _, pokemonProduct := range pokemonProducts {
     //converting a pokemonProduct to an array of strings
     record := []string{
       pokemonProduct.url,
       pokemonProduct.image,
       pokemonProduct.name,
       pokemonProduct.price,
     }

     //Writing a new CSV record
     writer.Write(record)
   }
   defer writer.Flush()
 }
