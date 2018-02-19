# web-crawler
Simple web crawler - Learning Go

First Go program that crawls a given domain name and that domain only. It does not follow external links.  Unit tests can be found in `crawler_test.go`

To run tests:
- Go to `src/` and run `go test`

To run tests with Coverage:
- Go to `src/` and run `go test -coverprofile=cover.out`. Use `go tool cover -func=cover.out` to see more details.

To run tests with HTML report:
- Go to `src/` and run `go tool cover -html=cover.out`


![Screen shot of visual testing](/img.png?raw=true "")


Bound by memory as `sync.Map` wouldn't be able to store the entire web on a local machine. If a network error occurs then solution should add link that was not queried correctly back onto the `queue`. 


//TO DO - re-evaluate design decision, opt for recusive approach (Breadth first search).
