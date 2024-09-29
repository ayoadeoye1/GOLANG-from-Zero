package main

import (
	"encoding/csv"
	"fmt"
	"os"

	// go colly package github
	"github.com/gocolly/colly"
)

type Product struct {
	name  string
	photo string
	price string
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.scrapingcourse.com"),
	)

	var products []Product

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Requesting:", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error Occurred:", err)
	})

	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		product := Product{}

		product.name = e.ChildText(".woocommerce-loop-product__title")
		product.photo = e.ChildAttr("a", "href")
		product.price = e.ChildText(".woocommerce-Price-amount.amount")
		
		products = append(products, product)
	})

	c.OnScraped(func(r *colly.Response) {
		file, err := os.Create("./scrape.csv")

		if err != nil {
			panic("Could not create the file")
		}
		defer file.Close()

		writer := csv.NewWriter(file)

		header := []string{"Name","Photo","Price"}

		writer.Write(header)

		for _, val := range products {
			body := []string{
				val.name,
				val.photo,
				val.price,
			}

			writer.Write(body)
		}
		defer writer.Flush()
	})

	c.Visit("https://www.scrapingcourse.com/ecommerce/")
}
