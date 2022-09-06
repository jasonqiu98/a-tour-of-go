/*
 1. the use of `sync.Mutex`
    lock & unlock for getters and setters
 2. the use of `var wg sync.WaitGroup`
    semaphore: `wg.Add(1)`, `wg.Done()`, `wg.Wait()`
*/
package main

import (
	"fmt"
	"sync"
)

type UrlResults struct {
	muBody  sync.Mutex
	muFound sync.Mutex
	body    map[string]string
	found   map[string]bool
}

func (res *UrlResults) SetBody(url string, bodyInfo string) {
	res.muBody.Lock()
	res.body[url] = bodyInfo
	res.muBody.Unlock()
}

func (res *UrlResults) GetBody(url string) string {
	res.muBody.Lock()
	defer res.muBody.Unlock()
	return res.body[url]
}

func (res *UrlResults) SetFound(url string, ifFound bool) {
	res.muFound.Lock()
	res.found[url] = ifFound
	res.muFound.Unlock()
}

func (res *UrlResults) GetFound(url string) bool {
	res.muFound.Lock()
	defer res.muFound.Unlock()
	return res.found[url]
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (res *UrlResults) Crawl(url string, depth int, fetcher Fetcher) {
	defer wg.Done()
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		res.SetFound(url, false)
		return
	}
	res.SetFound(url, true)
	res.SetBody(url, body)
	for _, u := range urls {
		wg.Add(1)
		go res.Crawl(u, depth-1, fetcher)
	}
	return
}

func (res *UrlResults) PrintUrlResults(url string) {
	defer wg.Done()
	if res.GetFound(url) {
		fmt.Printf("found: %s %q\n", url, res.GetBody(url))
	} else {
		fmt.Printf("not found: %s\n", url)
	}
}

// a global semaphore/wait group
var wg sync.WaitGroup

func main() {
	res := UrlResults{
		body:  make(map[string]string),
		found: make(map[string]bool),
	}

	wg.Add(1)
	go res.Crawl("https://golang.org/", 4, fetcher)
	wg.Wait()

	for url := range res.found {
		wg.Add(1)
		go res.PrintUrlResults(url)
	}
	wg.Wait()
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
