package main

import (
	"fmt"
	"sync"
	//"time"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

type UrlCrawled struct{
	urls map[string]bool
	mux sync.Mutex
}

func (u *UrlCrawled) has(url string) bool {
	u.mux.Lock()
	defer u.mux.Unlock()
	return u.urls[url]
}

func (u *UrlCrawled) add(url string) {
	u.mux.Lock()
	defer u.mux.Unlock()
	u.urls[url] = true
}

var urlCrawled = UrlCrawled{urls: make(map[string]bool)}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func CrawlRecursive(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	
	defer wg.Done()
	// 并行的抓取 URL。不重复抓取页面。
	if urlCrawled.has(url) {
		return
	}
	urlCrawled.add(url)

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
		go CrawlRecursive(u, depth-1, fetcher, wg)
	}
	return
}

func Crawl(url string, depth int, fetcher Fetcher) {
	var wg sync.WaitGroup
	wg.Add(1)
	go CrawlRecursive(url, depth, fetcher, &wg)
	wg.Wait()
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher 是返回若干结果的 Fetcher。
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

// fetcher 是填充后的 fakeFetcher。
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
