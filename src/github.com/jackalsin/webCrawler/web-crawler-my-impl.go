package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type fetchedMap struct {
	m map[string]bool
	sync.Mutex
}



// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	var fetched = fetchedMap{m: make(map[string]bool)}
	crawl(url, depth, fetcher, fetched)
}

func crawl(url string, depth int, fetcher Fetcher, fetched fetchedMap) {
	fmt.Printf("Begin to crawl %s\n", url)
	if depth <= 0 {
		return
	}

	fetched.Lock()
	if _, exist := fetched.m[url]; exist {
		fetched.Unlock() // already fetched.
		return
	}

	fetched.m[url] = true
	fetched.Unlock()

	// let's crawl
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println("Error in ", url, " is ", err)
		return
	}

	fmt.Printf("Found %s in url: %s \n", body, url)

	dones := make(chan bool)
	for _, childUrl := range urls {
		fmt.Printf("Going to crawl %s\n", childUrl)
		go func() {
			crawl(childUrl, depth - 1, fetcher, fetched)
			dones <- true
		}()
	}
	for i:= 0; i < len(urls); i++ {
		<- dones
	}
}
func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	// what is before semi-colon is executed before
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
