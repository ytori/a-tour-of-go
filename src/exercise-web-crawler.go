package main

import (
	"fmt"
	"sync"
)

type CachedFetcher struct {
	mux   sync.Mutex
	cache map[string]struct {
		body string
		urls []string
	}
	fetcher Fetcher
}

func NewCachedFetcher(f Fetcher) *CachedFetcher {
	return &CachedFetcher{cache: make(map[string]struct {
		body string
		urls []string
	}), fetcher: f}
}

func (f *CachedFetcher) Fetch(url string) (body string, urls []string, err error) {

	if val, ok := f.cache[url]; ok {
		return val.body, val.urls, nil
	}

	f.mux.Lock()
	defer f.mux.Unlock()

	if val, ok := f.cache[url]; ok {
		return val.body, val.urls, nil
	}

	fetchedBody, fetchedUrls, err := f.fetcher.Fetch(url)
	if err == nil {
		f.cache[url] = struct {
			body string
			urls []string
		}{body: fetchedBody, urls: fetchedUrls}
	}

	return fetchedBody, fetchedUrls, err
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go crawl(url, depth, fetcher, wg)
	wg.Wait()
}

func crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go crawl(u, depth-1, fetcher, wg)
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, NewCachedFetcher(fetcher))
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
