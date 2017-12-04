package main

import (
	"net/http"
	"fmt"
	"os"
	"golang.org/x/net/html"
	"strings"
	"net/url"
	"io"
	"sync"
	"time"
)

var baseUrl string
var visited sync.Map
var mutex sync.RWMutex //Using a mutex to prevent interleaving of fmt.Println() betweek different gorountines.

func main() {
	/*
		Prepare url for http request
	 */
	if len(os.Args) != 2 {
		baseUrl = "https://www.monzo.com/"
	} else {
		baseUrl = os.Args[1]
	}

	crawl(baseUrl)
}

func crawl(baseUrl string) {
	queue := make(chan string)
	finished := make(chan bool)

	fmt.Printf("Generating site map for %s ...", baseUrl)
	fmt.Println("")
	go processAll(queue, finished)
	queue <- baseUrl
	<-finished
	fmt.Println("\nComplete!")
}



/*
	Continuously loops through urls that have been added to channel 'queue'.
	If there is a long delay then break the loop and finish
 */
func processAll(queue chan string, finished chan bool) {
	for {
		select {
		case url := <-queue:
			if _, ok := visited.Load(url); ok {
				continue
			} else {
				visited.Store(url, true)
				go exploreUrl(url, queue)
			}
		case <-time.After(3 * time.Second):
			finished <- true
		}
	}
}
/*
	Scans current url and prints all links on the associated page. Does not follow external links.
	Every new 'monzo' url that is discovered is added to channel queue for processing.
 */
func exploreUrl(url string, queue chan string) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	mutex.Lock()
	fmt.Println("\nFetching ", url)
	defer resp.Body.Close()
	linksOnPage := getLinksOnPage(resp.Body)
	for _, link := range linksOnPage {
		if isExternalLink(link) {
			fmt.Println("\t" + link)
		} else {
			if url != "" {
				absolute := formatUrl(link, url)
				fmt.Println("\t" + absolute)
				if _, ok := visited.Load(absolute); !ok {
					queue <- absolute
				}
			}
		}
	}
	mutex.Unlock()
}

/*
	HTML Tokenizer - Parse html to extract tags
	ErrorToken	Error during tokenization (or end of document)
	StartTagToken	E.g. <a>
 */
func getLinksOnPage(body io.Reader) []string {
	allLinks := []string{}
	tokenizer := html.NewTokenizer(body)
	for {
		typeOfToken := tokenizer.Next()
		if typeOfToken == html.ErrorToken {
			return allLinks
		}
		token := tokenizer.Token()
		if typeOfToken == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					newLink := []string{ attr.Val }
					addOnce(&allLinks, newLink) //pass address of allLinks - able to modify orig var
				}
			}
		}
	}
}

func isExternalLink(url string) bool {
	isExternalLink :=
		strings.Index(url, "http") == 0 ||
			strings.Index(url, "/-play-store") == 0 || // redirects to external link
			strings.Index(url, "/docs") == 0 || // robots.txt restricts scan on /docs
			strings.Index(url, "tel:") == 0 ||
			strings.Contains(url, "#")
	return isExternalLink
}


/*
	Formats href to absolute path e.g. /about --> https://www.monzo.com/about
 */
func formatUrl(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return "Parse error"
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return "Parse error"
	}
	uri = baseUrl.ResolveReference(uri)
	return uri.String()
}

/*
	Adding new found link to list of all links if it doesn't already contain new link.
	Using reference to all links obj
 */
func addOnce(allLinks *[]string,  newLink []string) {
	for _, str := range newLink {
		if !contains(*allLinks, str) {
			*allLinks = append(*allLinks, str)
		}
	}
}

/*
	Checking to see new found link is not already contained in allLinks
 */
func contains(allLinks []string, newLink string) bool {
	var check bool
	for _, str := range allLinks {
		if str == newLink {
			check = true
			break
		}
	}
	return check
}
