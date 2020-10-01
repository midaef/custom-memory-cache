package cache

import "sync"

// Cache ...
type Cache struct {
	DataMutex *sync.Mutex
	HashSize  uint16
	Data      map[string]string
}

func NewCache() *Cache {
	return &Cache{
		DataMutex: new(sync.Mutex),
		HashSize:  1024,
	}
}
