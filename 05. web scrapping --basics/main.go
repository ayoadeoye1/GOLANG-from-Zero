package main

import (
	"fmt"

	// colly package for scraping
	"github.com/gocolly/colly"
)

type Product struct {
	photo string
	name  string
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

		product.photo = e.ChildAttr("a", "href")
		product.name = e.ChildText(".woocommerce-loop-product__title")
		product.price = e.ChildText(".woocommerce-Price-amount.amount")
		
		products = append(products, product)
	})

	c.Visit("https://www.scrapingcourse.com/ecommerce/")

	fmt.Println("Scraped Products:", products)
}
