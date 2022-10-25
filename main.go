package main

import "fmt"

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

func extractSiteMapURLs() {
	makeRequest()
}

// make request to a particular URL
func makeRequest() {

}

func scrapeURLs() {

}

func scrapePage() {

}

func crawlPage() {

}

func getSEOData() {

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
