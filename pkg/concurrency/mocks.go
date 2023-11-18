package concurrency

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type MockFetcher map[string]*mockResult

type mockResult struct {
	body string
	urls []string
}

var fetchSignal chan bool

// fetchSignalInstance is a singleton to access fetchSignal
func fetchSignalInstance() chan bool {
	if fetchSignal == nil {
		// Use buffered channel to avoid blocking
		fetchSignal = make(chan bool, 1000)
	}
	return fetchSignal
}

func (f MockFetcher) Fetch(url string) (string, []string, error) {
	fetchSignalInstance() <- true
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = MockFetcher{
	"http://golang.org/": &mockResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &mockResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &mockResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &mockResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

//types for produce and comsumers

type Stream struct {
	pos    int
	tweets []Tweet
}

type Tweet struct {
	Username string
	Text     string
}

// IsTalkingAboutGo is a mock process which pretend to be a sophisticated procedure to analyse whether tweet is talking about go or not
func (t *Tweet) IsTalkingAboutGo() bool {
	// simulate delay
	time.Sleep(330 * time.Millisecond)

	hasGolang := strings.Contains(strings.ToLower(t.Text), "golang")
	hasGopher := strings.Contains(strings.ToLower(t.Text), "gopher")

	return hasGolang || hasGopher
}

var mockdata = []Tweet{
	{
		"davecheney",
		"#golang top tip: if your unit tests import any other package you wrote, including themselves, they're not unit tests.",
	}, {
		"beertocode",
		"Backend developer, doing frontend featuring the eternal struggle of centering something. #coding",
	}, {
		"ironzeb",
		"Re: Popularity of Golang in China: My thinking nowadays is that it had a lot to do with this book and author https://github.com/astaxie/build-web-application-with-golang",
	}, {
		"beertocode",
		"Looking forward to the #gopher meetup in Hsinchu tonight with @ironzeb!",
	}, {
		"vampirewalk666",
		"I just wrote a golang slack bot! It reports the state of github repository. #Slack #golang",
	},
}

// ErrEOF returns on End of File error
var ErrEOF = errors.New("End of File")

// Next returns the next Tweet in the stream, returns EOF error if
// there are no more tweets
func (s *Stream) Next() (*Tweet, error) {

	// simulate delay
	time.Sleep(320 * time.Millisecond)
	if s.pos >= len(s.tweets) {
		return &Tweet{}, ErrEOF
	}

	tweet := s.tweets[s.pos]
	s.pos++

	return &tweet, nil
}
