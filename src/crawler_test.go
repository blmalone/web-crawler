package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
)

func TestIsExternalLink(t *testing.T) {
	status := 404
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
	}))
	defer ts.Close()


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
