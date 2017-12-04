package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
	"log"
	//"io/ioutil"
	"reflect"
)

var htmlWithOneLink = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8">
		<a href="https://www.google.com">Google</a></head><body></body></html>`

var htmlWithNoLinks = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"></head><body></body></html>`

func TestGetLinksOnPage(t *testing.T) {
	expectedLinksOnPage := []string {}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, htmlWithNoLinks)
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)

	if err != nil {
		log.Fatal(err)
	}

	actualLinksOnPage := getLinksOnPage(res.Body)

	if !reflect.DeepEqual(expectedLinksOnPage, actualLinksOnPage) {
		t.Errorf("getLinksOnPage(x) returns slice containing links")
	}

	expectedLinksOnPage = append(expectedLinksOnPage,`https://www.google.com`);
	ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, htmlWithOneLink)
	}))

	res, err = http.Get(ts.URL)

	if err != nil {
		log.Fatal(err)
	}

	actualLinksOnPage = getLinksOnPage(res.Body)

	if !reflect.DeepEqual(expectedLinksOnPage, actualLinksOnPage) {
		t.Errorf("getLinksOnPage(x) returns doesn't return expected link")
	}

}

func TestContains(t *testing.T) {
	allLinks := []string {"Blaine", "Bob", "Steve"}
	got := contains(allLinks, "Blaine")
	if !got {
		t.Errorf("contains(%q,%q) gives %v; want %v", allLinks, "Blaine", got, !got)
	}

	got = contains(allLinks, "Luke")
	if got {
		t.Errorf("contains(%q,%q) gives %v; want %v", allLinks, "Blaine", got, !got)
	}
}

func TestAddOnce(t *testing.T) {
	allLinks := []string {"Blaine", "Bob", "Steve"}
	target := "Blaine"
	newLink := []string{ target }
	addOnce(&allLinks, newLink)
	count := 0
	for _, link := range allLinks {
		if link == target {
			count++
		}
	}
	if count > 1 {
		t.Errorf("addOnce(%q,%q) adds %v %v times", allLinks, target, target, count)
	}
}

func TestFormatUrl(t *testing.T) {
	expectedResult := "https://www.google.co.uk/maps"
	result := formatUrl("/maps", "https://www.google.co.uk")

	if result != expectedResult {
		t.Errorf("The url was formatted incorrectly")
	}
}
