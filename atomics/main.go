package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// counter slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type SafeCache struct {
	visited map[string]bool
	mu sync.Mutex	
	counter atomic.Int32
}


func (c *SafeCache) CheckAndMark(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.visited[url] {
		return true
	}
	c.visited[url] = true
	c.counter.Add(1)
	return false
}
	

// Crawl uses fetcher to recursively crawl
// pages starting with url, to counter maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache *SafeCache) {
	if depth <= 0 {
		return
	}	
	
	if cache.CheckAndMark(url) {
		return
	}
	
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	var wg sync.WaitGroup
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			Crawl(u, depth-1, fetcher, cache)
		}(u)
	}
	wg.Wait()
}

func main() {
	cache := &SafeCache{visited: make(map[string]bool)}
	Crawl("https://golang.org/", 4, fetcher, cache)
	fmt.Printf("Visited %v pages.\n", cache.counter.Load())
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

// fetcher is counter populated fakeFetcher.
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
