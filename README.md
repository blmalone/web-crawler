# web-crawler
Simple web crawler - Learning Go

First Go program that crawls a given domain name and that domain only. It does not follow external links.  Unit tests can be found in `crawler_test.go`

To run tests:
- Go to `src/` and run `go test`

To run tests with Coverage:
- Go to `src/` and run `go test -coverprofile=cover.out`. Use `go tool cover -func=cover.out` to see more details.
