package racing

import (
	"strconv"
	"sync"
	"time"
)

type MockDB struct{}

// Get only returns an empty string, as this is only for demonstration purposes
func (*MockDB) Get(key string) (string, error) {
	d, _ := time.ParseDuration("20ms")
	time.Sleep(d)
	return "", nil
}

// GetMockDB returns an instance of MockDB
func GetMockDB() *MockDB {
	return &MockDB{}
}

const (
	cycles        = 15
	callsPerCycle = 100
)

// RunMockServer simulates a running server, which accesses the
// key-value database through our cache
func RunMockServer(cache *KeyStoreCache) {
	var wg sync.WaitGroup

	for c := 0; c < cycles; c++ {
		wg.Add(1)
		go func() {
			for i := 0; i < callsPerCycle; i++ {

				cache.Get("Test" + strconv.Itoa(i))

			}
			wg.Done()
		}()
	}

	wg.Wait()
}
