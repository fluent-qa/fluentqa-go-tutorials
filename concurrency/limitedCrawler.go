package concurrency

import (
	"fmt"
	"sync"
)

func Crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	fmt.Println(len(urls))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	wg.Add(len(urls))
	for _, u := range urls {
		// Do not remove the `go` keyword, as Crawl() must be
		// called concurrently
		fmt.Printf("start to fetch %s", u)
		go Crawl(u, depth-1, wg)
	}

}

func Run() {
	var wg sync.WaitGroup
	wg.Add(1)
	Crawl("http://golang.org", 4, &wg)
	wg.Wait()
}

// ## Keypoint:
// 1. golang sync WaitGroup
// 2. How to work as Wait group
