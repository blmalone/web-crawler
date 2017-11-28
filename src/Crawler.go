package main

import (
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
	"os"
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

	fmt.Println(string(body))

	/*
		Make http GET request
 	*/
	log.Printf("Fetching %v\n", url)
	resp, err = http.Get(url)

	if err != nil {
		log.Fatal("Error GET request.")
	}
	defer resp.Body.Close() // Close body when read to completion - declared here fro readability
	body, err = ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	fmt.Println("Hello Blaine")
}

func prepareRobotsUrl(url string) string {
	if url[len(url)-1] != '/' {
		url = url + "/"
	}
	return  url + "robots.txt"
}
