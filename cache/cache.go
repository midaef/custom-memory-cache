package cache

import (
	"errors"
	"sync"
	"time"

	"github.com/midaef/custom-memory-cache/models"
)

var (
	ErrKeyNotFound     = errors.New("custom-memory-cache: Key not found")
	ErrCacheIsNotAlive = errors.New("custom-memory-cache: Cache is not alive")
	ErrCacheIsEmpty    = errors.New("custom-memory-cache: Cache is empty")
)

// Cache ...
type Cache struct {
	DataMutex  *sync.Mutex
	Size       uint16
	GCInterval time.Duration
	Data       map[string]models.Data
}

func NewCache(size uint16, gcInterval time.Duration) *Cache {
	cache := &Cache{
		DataMutex:  new(sync.Mutex),
		Data:       make(map[string]models.Data, size),
		Size:       size,
		GCInterval: gcInterval,
	}
	if gcInterval > 0 {
		gc := NewGC(cache)
		gc.Start()
	}
	return cache
}

func (cache *Cache) Write(key string, value interface{}, lifeTime time.Duration) {
	var lfTime int64
	if lifeTime > 0 {
		lfTime = time.Now().Add(lifeTime).UnixNano()
	}
	cache.DataMutex.Lock()
	cache.Data[key] = models.Data{
		Data:     value,
		LifeTime: lfTime,
		Created:  time.Now(),
		Alive:    true,
	}
	cache.DataMutex.Unlock()
}

func (cache *Cache) Read(key string) (interface{}, error) {
	cache.DataMutex.Lock()
	defer cache.DataMutex.Unlock()
	data, found := cache.Data[key]
	if !found {
		return nil, ErrKeyNotFound
	}
	if data.LifeTime > 0 {
		if time.Now().UnixNano() > data.LifeTime {
			return nil, ErrCacheIsNotAlive
		}
	}
	return data.Data, nil
}

func (cache *Cache) ReadAll() (map[string]models.Data, error) {
	cache.DataMutex.Lock()
	defer cache.DataMutex.Unlock()
	if cache.Data != nil {
		return cache.Data, nil
	}
	return nil, ErrCacheIsEmpty
}

func (cache *Cache) Delete(key string) error {
	cache.DataMutex.Lock()
	defer cache.DataMutex.Unlock()
	_, ok := cache.Data[key]
	if !ok {
		return ErrKeyNotFound
	}
	delete(cache.Data, key)
	return nil
}

func (cache *Cache) DeleteAll() error {
	cache.DataMutex.Lock()
	defer cache.DataMutex.Unlock()
	if cache.Data != nil {
		for key := range cache.Data {
			delete(cache.Data, key)
		}
		return nil
	}
	return ErrCacheIsEmpty
}
