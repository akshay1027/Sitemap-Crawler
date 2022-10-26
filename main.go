package main

import (
	"fmt"
	"log"
)

type SeoData struct {
	URL             string
	Title           string
	H1              string
	MetaDescription string
	StatusCode      int
}

type Parser interface {
}

type DefaultParser struct {
}

// browser telling the website that a human being is
// scraping the website instead of bots etc. its a signature to valid human
type userAgents struct {
}

func randomUserAgent() {

}

// Flow:
// 1. get the start url and add it to WorkList
// 2. copy the WorkList to a new variable list, to iterate on them
// 3. range over the start url, get the response
// 4. from response, extract urls, check if these urls are sitemap
// 5. if so append them to toCrwal(array of string)
// 6. return the list of urls to crawl

func extractSiteMapURLs(startURL string) []string {
	// WorkList is a channel which will store all the URLs that needs to be scraped
	WorkList := make(chan []string)

	// intialise empty array of strings, store all the urls present from response of crawling startUrl
	toCrawl := []string{}

	go func() {
		// add starting url to WorkList channel
		WorkList <- []string{startURL}
	}()

	// move urls from worklkist to list, we are going to use this to iterate over the urls
	list := <-WorkList

	for n := 1; n > 0; n-- {
		for _, link := range list {
			// create infinite loop scneario to add all the urls toCrawl array
			n++

			// go routine
			go func(link string) {
				res, err := makeRequest(link)
				if err != nil {
					log.Println("Error retriving url: %s", link)
				}

				urls, _ := extractURLs(res)
				if err != nil {
					log.Println("Error retriving url: %s", link)
				}

				sitemapFiles, pages := isSitemap(urls)
				if sitemapFiles != nil {
					WorkList <- sitemapFiles
				}

				for _, page := range pages {
					toCrawl = append(toCrawl, page)
				}
			}(link)
		}
	}

	return toCrawl
}

// make request to a particular URL
func makeRequest(url string) string {

}

func scrapeURLs() {

}

func scrapePage() {

}

func crawlPage() {

}

func getSEOData() {

}

func extractURLs() {

}

// scrape url and return slice of data (type SeoData)
func scrapeSiteMap(url string) []SeoData {
	results := extractSiteMapURLs(url)
	res := scrapeURLs(results)

	return res
}

func main() {
	p := DefaultParser{}
	results := scrapeSiteMap("")

	for _, result := range results {
		fmt.Println(result)
	}
}
