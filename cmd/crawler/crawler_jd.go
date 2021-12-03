package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

func main() {
	// Array containing all the known URLs in a sitemap
	knownUrls := []string{}

	// Create a Collector specifically for Shopify
	c := colly.NewCollector(colly.AllowedDomains("item.jd.com"))
	extensions.RandomUserAgent(c)
    extensions.Referer(c)
	c.OnHTML("div.sku-name",func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
		
	})
	c.OnHTML("span.p-price",func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})
	c.OnHTML("div#choose-attrs",func(h *colly.HTMLElement) {
		fmt.Println(h.DOM.Html())
	})
	// Start the collector
	c.Visit("https://item.jd.com/100026667928.html")
	
	fmt.Println("All known URLs:")
	for _, url := range knownUrls {
		fmt.Println("\t", url)
	}
	fmt.Println("Collected", len(knownUrls), "URLs")

	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var example string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://item.jd.com/100026667928.html`),
		// wait for footer element is visible (ie, page is loaded)
		//chromedp.WaitVisible(`body > footer`),
		// find and click "Expand All" link
		//chromedp.Click(`#pkg-examples > div`, chromedp.NodeVisible),
		// retrieve the value of the textarea
		//chromedp.Value(`#example_After .play .input textarea`, &example),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Go's time.After example:\n%s", example)
}