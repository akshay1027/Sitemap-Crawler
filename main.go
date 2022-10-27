package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
	// "net/http"
	// "github.com/PuerkitoBio/goquery"
)

type SeoData struct {
	URL             string
	Title           string
	H1              string
	MetaDescription string
	StatusCode      int
}

// type Parser interface {
// }

// type DefaultParser struct {
// }

// browser telling the website that a human being is
// scraping the website instead of bots etc. its a signature to valid human
var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
}

func randomUserAgent() string {
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}

func isSitemap(urls []string) ([]string, []string) {
	sitemapFiles := []string{}
	pages := []string{}
	for _, page := range urls {
		foundSitemap := strings.Contains(page, "xml")
		if foundSitemap == true {
			fmt.Println("Found Sitemap", page)
			sitemapFiles = append(sitemapFiles, page)
		} else {
			pages = append(pages, page)
		}
	}
	return sitemapFiles, pages
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

func extractURLs(url string) SeoData {

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
