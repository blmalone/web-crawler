package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
	"log"
	"reflect"
)

var htmlWithOneLink = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8">
		<a href="https://www.google.com">Google</a></head><body></body></html>`

var htmlWithNoLinks = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"></head><body></body></html>`

var htmlWithTwoLinks = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8">
		<a href="https://www.google.com">Google</a><a href="/about">About</a></head><body></body></html>`

var htmlWithDuplicateLinks = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8">
		<a href="https://www.google.com">Google</a><a href="/about">About</a>
		<a href="/about">About</a><a href="/about">About</a>
		</head><body></body></html>`

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

	expectedLinksOnPage = append(expectedLinksOnPage,`/about`);
	ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, htmlWithTwoLinks)
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

	result = formatUrl("", ":")

	if result != "Parse error" {
		t.Errorf("Invalid url shouldn't have been parsed")
	}

	result = formatUrl(":", "")

	if result != "Parse error" {
		t.Errorf("Invalid url shouldn't have been parsed")
	}
}

func TestIsExternalLink(t *testing.T) {
	externalLink := "https://www.apple.com/"

	internalLink := "monzo.com/blog/"

	result := isExternalLink(externalLink)

	if !result {
		t.Errorf("isExternalLink(%q) returned false when %q is an external link",
			externalLink, externalLink)
	}

	result = isExternalLink(internalLink)

	if result {
		t.Errorf("isExternalLink(%q) returned true when %q isn't an external link",
			internalLink, internalLink)
	}
}

func TestNoDeadLockInCrawl(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, htmlWithDuplicateLinks)
	}))
	defer ts.Close()
	crawl(ts.URL)
}
