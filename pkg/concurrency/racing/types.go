package racing

import "container/list"

// CacheSize determines how big the cache can grow
const CacheSize = 100

// KeyStoreCacheLoader is an interface for the KeyStoreCache
type KeyStoreCacheLoader interface {
	// Load implements a function where the cache should gets it's content from
	Load(string) string
}

type Page struct {
	Key   string
	Value string
}

type KeyStoreCache struct {
	Cache map[string]*list.Element
	Pages list.List
	Load  func(string) string
}

// New creates a new KeyStoreCache
func New(load KeyStoreCacheLoader) *KeyStoreCache {
	return &KeyStoreCache{
		Load:  load.Load,
		Cache: make(map[string]*list.Element),
	}
}

// Get gets the key from cache, loads it from the source if needed
func (k *KeyStoreCache) Get(key string) string {
	if e, ok := k.Cache[key]; ok {
		k.Pages.MoveToFront(e)
		return e.Value.(Page).Value
	}
	// Miss - load from database and save it in cache
	p := Page{key, k.Load(key)}
	// if cache is full remove the least used item
	if len(k.Cache) >= CacheSize {
		end := k.Pages.Back()
		// remove from map
		delete(k.Cache, end.Value.(Page).Key)
		// remove from list
		k.Pages.Remove(end)
	}
	k.Pages.PushFront(p)
	k.Cache[key] = k.Pages.Front()
	return p.Value
}

// Loader implements KeyStoreLoader
type Loader struct {
	DB *MockDB
}

// Load gets the data from the database
func (l *Loader) Load(key string) string {
	val, err := l.DB.Get(key)
	if err != nil {
		panic(err)
	}

	return val
}

func RunRacing() *KeyStoreCache {
	loader := Loader{
		DB: GetMockDB(),
	}
	cache := New(&loader)

	RunMockServer(cache)

	return cache
}
