package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"os"
	"golang.org/x/net/html"
	"strings"
)



func main() {
	/*
		Prepare url for http request
	 */
	var url string
	if len(os.Args) != 2 {
		url = "https://www.monzo.com/"
	} else {
		url = os.Args[1]
	}
	robotsFileUrl := prepareRobotsUrl(url)
	fmt.Println("URL to Crawl: ", url);


	/*
		Make http GET request
 	*/
	resp, err := http.Get(robotsFileUrl)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body) + "\n")

	/*
		Make http GET request
 	*/
	resp, err = http.Get(url)

	if err != nil {
		fmt.Println("ERROR: failed to GET data from ", url)
	}

	defer resp.Body.Close() // Close body when read to completion - declared here for readability

	/*
		HTML Tokenizer - Parse html to extract tags
		ErrorToken	Error during tokenization (or end of document)
		StartTagToken	E.g. <a>
	 */
	tokenizer := html.NewTokenizer(resp.Body)
	for {
		nextToken := tokenizer.Next()
		switch nextToken {
		case html.ErrorToken: // End of the document, we're done
			return
		case html.StartTagToken:
			currentToken := tokenizer.Token()
			isAnchor := currentToken.Data == "a"
			if !isAnchor { //Fail fast
				continue
			}
			ok, url := getHrefAttribute(currentToken) //Returns ok as 'true' if attribute found
			if !ok {
				continue
			}
			isExternalHttpLink := strings.Index(url, "http") == 0
			if isExternalHttpLink {
				fmt.Println(url)
			}
		}
	}

}

func getHrefAttribute(t html.Token) (ok bool, href string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}
	return
}

func prepareRobotsUrl(url string) string {
	if url[len(url) - 1] != '/' {
		url = url + "/"
	}
	return url + "robots.txt"
}
